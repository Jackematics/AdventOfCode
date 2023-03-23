import Cargo from './cargo';
import CrateMover from './enums/crate-mover';
import FileProcessor from './file-processor';

const { cargo, moveOrders } =
  FileProcessor.processInputIntoCargoData('src/input.txt');

// const mover = CrateMover.NineThousand;
const mover = CrateMover.NineThousandAndOne;

cargo.executeMoveOrders(moveOrders, mover);

const result: string = cargo.getTopCrates();

console.log('Part 1 result: ', result);
