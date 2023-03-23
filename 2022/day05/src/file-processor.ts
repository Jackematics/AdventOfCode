import * as fs from 'fs';
import Cargo from './cargo';
import CargoData from './types/cargo-data';

export default class FileProcessor {
  public static processInputIntoCargoData(inputPath: string): CargoData {
    const raw = fs.readFileSync(inputPath).toString().split('\n');
    const dividerIndex = raw.indexOf('');
    const cargo = this.processCargo(raw.slice(0, dividerIndex - 1).reverse());
    const moveOrders = this.processMoveOrders(raw.slice(dividerIndex + 1));

    return {
      cargo,
      moveOrders,
    };
  }

  private static processCargo(rawCargo: string[]): Cargo {
    const colCount = rawCargo[0].split('[').length - 1;
    const cargoArray = rawCargo
      .reduce(this.formatIntoColumnsReducer, new Array(colCount).fill([]))
      .map((col) => col.filter((crate) => crate !== ' '));

    return new Cargo(cargoArray);
  }

  private static formatIntoColumnsReducer(
    acc: string[][],
    row: string
  ): string[][] {
    const gapSize = 4;
    const firstIndexOffset = 1;
    return acc.map((col, index) => [
      ...col,
      row[gapSize * index + firstIndexOffset],
    ]);
  }

  private static processMoveOrders(rawMoveOrders: string[]) {
    return rawMoveOrders.map((rawMoveOrder) => {
      const split = rawMoveOrder.split(' ');
      const quantityIndex = 1;
      const fromIndex = 3;
      const toIndex = 5;

      return {
        quantity: parseInt(split[quantityIndex]),
        from: parseInt(split[fromIndex]) - 1,
        to: parseInt(split[toIndex]) - 1,
      };
    });
  }
}
