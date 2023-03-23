import MoveOrder from './types/move-order';

export default class TestData {
  public static readonly exampleCargoArray = [
    ['Z', 'N'],
    ['M', 'C', 'D'],
    ['P'],
  ];
  public static readonly moveOrders: MoveOrder[] = [
    { quantity: 1, from: 1, to: 0 },
    { quantity: 3, from: 0, to: 2 },
    { quantity: 2, from: 1, to: 0 },
    { quantity: 1, from: 0, to: 1 },
  ];
}
