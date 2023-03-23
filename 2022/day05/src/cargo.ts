import CrateMover from './enums/crate-mover';
import MoveOrder from './types/move-order';

export default class Cargo {
  public readonly cargo: string[][];

  constructor(cargo: string[][]) {
    this.cargo = [...cargo];
  }

  public getCargo() {
    return this.cargo;
  }

  public getTopCrates() {
    return this.cargo.reduce(this.getTopCrateReducer, '');
  }

  private getTopCrateReducer(acc: string, stack: string[]) {
    return `${acc}${stack[stack.length - 1] ?? ' '}`;
  }

  public executeMoveOrders(moveOrders: MoveOrder[], mover: CrateMover): void {
    moveOrders.map((moveOrder) => {
      this.executeMoveOrder(moveOrder, mover);
    });
  }

  private executeMoveOrder(moveOrder: MoveOrder, mover: CrateMover): void {
    const inTransit: string[] = this.cargo[moveOrder.from].splice(
      -moveOrder.quantity,
      moveOrder.quantity
    );
    const transitType =
      mover === CrateMover.NineThousand ? inTransit.reverse() : inTransit;

    this.cargo[moveOrder.to].push(...transitType);
  }
}
