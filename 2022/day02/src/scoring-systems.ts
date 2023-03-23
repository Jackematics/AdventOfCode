export type ScoringSystem = { [key: string]: number };

export default class ScoringSystems {
  public static readonly part1ScoringSystem: ScoringSystem = {
    ['A X']: 4,
    ['A Y']: 8,
    ['A Z']: 3,
    ['B X']: 1,
    ['B Y']: 5,
    ['B Z']: 9,
    ['C X']: 7,
    ['C Y']: 2,
    ['C Z']: 6,
  };

  public static readonly part2ScoringSystem: ScoringSystem = {
    ['A X']: 3,
    ['A Y']: 4,
    ['A Z']: 8,
    ['B X']: 1,
    ['B Y']: 5,
    ['B Z']: 9,
    ['C X']: 2,
    ['C Y']: 6,
    ['C Z']: 7,
  };
}
