import Cargo from './cargo';
import CrateMover from './enums/crate-mover';
import TestData from './test-data';

describe('SupplyStackRearranger()', () => {
  it.each([
    {
      cargo: new Cargo([['Z', 'N'], ['M', 'C', 'D'], ['P']]),
      movementOrders: TestData.moveOrders.slice(0, 1),
      expected: 'DCP',
    },
    {
      cargo: new Cargo([['Z', 'N', 'D'], ['M', 'C'], ['P']]),
      movementOrders: TestData.moveOrders.slice(1, 2),
      expected: ' CZ',
    },
    {
      cargo: new Cargo([[], ['M', 'C'], ['P', 'D', 'N', 'Z']]),
      movementOrders: TestData.moveOrders.slice(2, 3),
      expected: 'M Z',
    },
    {
      cargo: new Cargo([['C', 'M'], [], ['P', 'D', 'N', 'Z']]),
      movementOrders: TestData.moveOrders.slice(3, 4),
      expected: 'CMZ',
    },
  ])(
    `top crates should be $expected after input moves are executed using CrateMover 9000`,
    ({ cargo, movementOrders, expected }) => {
      cargo.executeMoveOrders(movementOrders, CrateMover.NineThousand);
      expect(cargo.getTopCrates()).toBe(expected);
    }
  );

  it('should return the correct top crates after a sequence of moves are executed using CrateMover 9000', () => {
    const cargo = new Cargo([['Z', 'N'], ['M', 'C', 'D'], ['P']]);
    cargo.executeMoveOrders(TestData.moveOrders, CrateMover.NineThousand);
    expect(cargo.getTopCrates()).toBe('CMZ');
  });

  it('should return the correct top crates after a sequence of moves are executed using CrateMover 9001', () => {
    const cargo = new Cargo([['Z', 'N'], ['M', 'C', 'D'], ['P']]);
    cargo.executeMoveOrders(TestData.moveOrders, CrateMover.NineThousandAndOne);
    expect(cargo.getTopCrates()).toBe('MCD');
  });
});
