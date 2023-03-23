import FileProcessor from './file-processor';
import TestData from './test-data';
import CargoData from './types/cargo-data';

describe('FileProcessor()', () => {
  it('should convert a text file of lines into an array of number arrays', () => {
    const result: CargoData = FileProcessor.processInputIntoCargoData(
      'src/example-input.txt'
    );
    const cargoExpected = TestData.exampleCargoArray;
    const expectedMoveOrders = TestData.moveOrders;

    expect(result.cargo.getCargo()).toStrictEqual(cargoExpected);
    expect(result.moveOrders).toStrictEqual(expectedMoveOrders);
  });
});
