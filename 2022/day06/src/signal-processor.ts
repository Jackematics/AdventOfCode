export default class SignalProcessor {
  public static findStartOfPacketMarkerIndex(packet: string): number {
    return this.findMarker(packet, 4);
  }

  public static findStartOfMessageMarkerIndex(packet: string): number {
    return this.findMarker(packet, 14);
  }

  private static findMarker(
    packet: string,
    markerSequenceLength: number
  ): number {
    return (
      packet
        .split('')
        .map((char, index) => packet.slice(index, index + markerSequenceLength))
        .slice(0, -(markerSequenceLength + 1))
        .findIndex((potentialDistinct) =>
          this.allCharsDistinct(potentialDistinct)
        ) + markerSequenceLength
    );
  }

  private static allCharsDistinct(string: string): boolean {
    return string.split('').length === new Set(string).size;
  }
}
