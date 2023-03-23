import { Part } from './enum';
import FileProcessor from './file-processor';
import Monkey from './monkey';
import MonkeyBusinessAnalyser from './monkey-analyser';
import MonkeyModeller from './monkey-modeller';
import { MonkeyState } from './types';

describe('Monkey Business', () => {
  describe('Monkey', () => {
    it('should count 1 time for 1 monkey inspecting 1 item', () => {
      const startingState: MonkeyState = {
        items: [79],
        divisor: 23,
        operation: (oldWorry: number) => oldWorry * 19,
        test: (worry: number) => (worry % 19 === 0 ? 2 : 0),
      };

      const monkey = new Monkey(startingState);

      monkey.part1TakeTurn();

      expect(monkey.itemInspectionCount).toBe(1);
    });

    it('should apply worry operation onto items and inspection count be number of items', () => {
      const startingState: MonkeyState = {
        items: [79, 51],
        divisor: 23,
        operation: (oldWorry: number) => oldWorry * 19,
        test: (worry: number) => (worry % 19 === 0 ? 2 : 0),
      };

      const monkey = new Monkey(startingState);

      monkey.part1TakeTurn();

      expect(monkey.itemInspectionCount).toBe(2);
      expect(monkey.items).toStrictEqual([500, 323]);
    });

    it('should throw items after inspection to the correct monkeys', () => {
      const startingState: MonkeyState = {
        items: [79, 51],
        divisor: 23,
        operation: (oldWorry: number) => oldWorry * 19,
        test: (worry: number) => (worry % 23 === 0 ? 2 : 3),
      };

      const monkey = new Monkey(startingState);

      monkey.part1TakeTurn();
      const expectedThrownItems = [
        { monkey: 3, item: 500 },
        { monkey: 3, item: 323 },
      ];

      expect(monkey.throwItems()).toStrictEqual(expectedThrownItems);
      expect(monkey.items).toStrictEqual([]);
    });

    it('should receive an item from another monkey', () => {
      const startingState: MonkeyState = {
        items: [79, 51],
        divisor: 23,
        operation: (oldWorry: number) => oldWorry * 19,
        test: (worry: number) => (worry % 23 === 0 ? 2 : 3),
      };

      const monkey = new Monkey(startingState);

      monkey.receiveItem(23);

      expect(monkey.items).toStrictEqual([79, 51, 23]);
    });
  });

  describe('MonkeyModeller', () => {
    it('should correctly model a group of monkeys throwing items', () => {
      const monkeys = FileProcessor.processInput('src/example-input.txt');

      MonkeyModeller.playRounds(monkeys, 20, Part.One);

      expect(MonkeyBusinessAnalyser.calculateMonkeyBusinessLevel(monkeys)).toBe(
        10605
      );
    });

    it('should correctly model a group of monkeys throwing items', () => {
      const monkeys = FileProcessor.processInput('src/example-input.txt');

      MonkeyModeller.playRounds(monkeys, 10000, Part.Two);

      expect(MonkeyBusinessAnalyser.calculateMonkeyBusinessLevel(monkeys)).toBe(
        2713310158
      );
    });
  });
});
