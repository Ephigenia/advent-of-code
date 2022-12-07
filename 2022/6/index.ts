import { readFile } from "fs/promises";
import { resolve } from "path";
import { EventEmitter } from "stream";

interface SignalStartMarkerEvent {
  pos: number;
  marker: string;
}

class ElvDetector extends EventEmitter {
  buffer: string[] = [];
  startOfSignalDetected: number|undefined;

  readFromString(input: string) {
    for (let i = 0; i < input.length; i++) {
      const char = input.substring(i, i + 1);
      this.write(char);
    }
  }

  last(n: number) {
    return this.buffer.slice(-n);
  }

  isMarker(bytes: string[], length: number) {
    return (new Set(bytes)).size === length;
  }

  signalMarkerLength = 4;
  isSignalStartMarker(bytes: string[]) {
    return this.isMarker(bytes, this.signalMarkerLength);
  }

  checkForSignalMarker() {
    if (
      !this.startOfSignalDetected &&
      this.buffer.length >= this.signalMarkerLength &&
      this.isSignalStartMarker(this.last(this.signalMarkerLength))
    ) {
      this.emit(ElvDetector.EVENT_START_MARKER, {
        pos: this.buffer.length,
        marker: this.last(this.signalMarkerLength).join(''),
      });
      this.startOfSignalDetected = this.buffer.length;
    }
  }

  messageMarkerLength = 14;
  isMessageMarker(bytes: string[]) {
    return this.isMarker(bytes, this.messageMarkerLength);
  }

  static EVENT_MESSAGE_MARKER = 'messageMarker';
  static EVENT_START_MARKER = 'signalStartMarker';

  checkForMessageMarker() {
    if (
      this.buffer.length >= this.messageMarkerLength &&
      this.isMessageMarker(this.last(this.messageMarkerLength))
    ) {
      this.emit(ElvDetector.EVENT_MESSAGE_MARKER, {
        pos: this.buffer.length,
        marker: this.last(this.messageMarkerLength).join(''),
      });
      this.startOfSignalDetected = this.buffer.length;
    }
  }

  write(byte: string) {
    this.buffer.push(byte);
    this.checkForSignalMarker();
    this.checkForMessageMarker();
  }

  toString() {
    return this.buffer.join('');
  }
}

async function main(filename: string) {
  const inputFilename = resolve(__dirname, filename);
  const raw = (await readFile(inputFilename)).toString().trim();

  const detector = new ElvDetector();
  detector.on(ElvDetector.EVENT_MESSAGE_MARKER, function(event: SignalStartMarkerEvent) {
    console.log('message marker found %d (%j)', event.pos, event.marker);
  });
  detector.on(ElvDetector.EVENT_START_MARKER, function(event: SignalStartMarkerEvent) {
    console.log('signal marker found at %d (%j)', event.pos, event.marker);
  });
  detector.readFromString(raw);
}

main('input.txt');
// main('inputTest.txt');
