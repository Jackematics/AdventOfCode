import { MonkeyState, ThrownItem } from './types';

export default class Monkey {
  private _items: number[];
  private _divisor: number;
  private _itemInspectionCount: number = 0;

  private readonly operation: (oldWorry: number) => number;
  private readonly test: (worry: number) => number;

  constructor(state: MonkeyState) {
    this._items = state.items;
    this._divisor = state.divisor;
    this.operation = state.operation;
    this.test = state.test;
  }

  get items() {
    return this._items;
  }

  get divisor() {
    return this._divisor;
  }

  get itemInspectionCount() {
    return this._itemInspectionCount;
  }

  public part1TakeTurn() {
    this._items = this._items.map((item: number) => {
      this._itemInspectionCount++;
      return Math.floor(this.operation(item) / 3);
    });
  }

  public part2TakeTurn(modulus: number) {
    this._items = this._items.map((item: number) => {
      this._itemInspectionCount++;
      return this.operation(item % modulus);
    });
  }

  private lowestCommonMultiple() {
    const gcd = (a: number, b: number): number => (a ? gcd(b % a, a) : b);
    const lcm = (a: number, b: number): number => (a * b) / gcd(a, b);

    return this._items.reduce(lcm);
  }

  public throwItems(): ThrownItem[] {
    const toThrow = this._items.map((item) =>
      this.calculateToThrowLocation(item)
    );

    this._items = [];
    return toThrow;
  }

  private calculateToThrowLocation(item: number) {
    return {
      monkey: this.test(item),
      item,
    };
  }

  public receiveItem(item: number): void {
    this._items.push(item);
  }
}
