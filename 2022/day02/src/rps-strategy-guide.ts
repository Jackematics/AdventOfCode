import { ScoringSystem } from './scoring-systems';

export default class RpsStrategyGuide {
  public static calculateTotalScore(
    games: string[],
    scoringSystem: ScoringSystem
  ): number {
    return games.reduce(
      (acc: number, element: string) => acc + scoringSystem[element],
      0
    );
  }
}
