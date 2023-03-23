export default class PriorityCalculator {
  public static calculatePriority(itemType: string) {
    this.validate(itemType);

    const asciiUppercaseOffset = 38;
    const asciiLowercaseOffset = 96;

    return itemType.toUpperCase() === itemType
      ? itemType.charCodeAt(0) - asciiUppercaseOffset
      : itemType.charCodeAt(0) - asciiLowercaseOffset;
  }

  private static validate(itemType: string): void {
    this.errorIfNotChar(itemType);
    this.errorIfInvalidAscii(itemType.charCodeAt(0));
  }

  private static errorIfNotChar(itemType: string): void {
    if (itemType.length !== 1) throw new Error('item type must be a character');
  }

  private static errorIfInvalidAscii(charcode: number): void {
    if (charcode < 64 || (90 < charcode && charcode < 97) || charcode > 122)
      throw new Error('item have a valid ascii value (a-z or A-Z)');
  }
}
