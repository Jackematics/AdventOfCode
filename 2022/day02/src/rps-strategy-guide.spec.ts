import RpsStrategyGuide from './rps-strategy-guide';
import ScoringSystems from './scoring-systems';

describe('RpsStrategyGuide()', () => {
  it.each([
    { input: ['A X'], result: 4 },
    { input: ['A Y'], result: 8 },
    { input: ['A Z'], result: 3 },
    { input: ['B X'], result: 1 },
    { input: ['B Y'], result: 5 },
    { input: ['B Z'], result: 9 },
    { input: ['C X'], result: 7 },
    { input: ['C Y'], result: 2 },
    { input: ['C Z'], result: 6 },
  ])(
    'should return the score $result for game $input and part 1 scoring system',
    ({ input, result }) => {
      expect(
        RpsStrategyGuide.calculateTotalScore(
          input,
          ScoringSystems.part1ScoringSystem
        )
      ).toBe(result);
    }
  );

  it('should calculate the part 1 total score', () => {
    const input = ['A Y', 'B X', 'C Z'];
    expect(
      RpsStrategyGuide.calculateTotalScore(
        input,
        ScoringSystems.part1ScoringSystem
      )
    ).toBe(15);
  });

  it.each([
    { input: ['A X'], result: 3 },
    { input: ['A Y'], result: 4 },
    { input: ['A Z'], result: 8 },
    { input: ['B X'], result: 1 },
    { input: ['B Y'], result: 5 },
    { input: ['B Z'], result: 9 },
    { input: ['C X'], result: 2 },
    { input: ['C Y'], result: 6 },
    { input: ['C Z'], result: 7 },
  ])(
    'should return the score $result for game $input and part 2 scoring system',
    ({ input, result }) => {
      expect(
        RpsStrategyGuide.calculateTotalScore(
          input,
          ScoringSystems.part2ScoringSystem
        )
      ).toBe(result);
    }
  );

  it('should calculate the part 2 total score', () => {
    const input = ['A Y', 'B X', 'C Z'];
    expect(
      RpsStrategyGuide.calculateTotalScore(
        input,
        ScoringSystems.part2ScoringSystem
      )
    ).toBe(12);
  });
});
