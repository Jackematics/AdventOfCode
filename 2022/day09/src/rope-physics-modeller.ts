import Motion from './motion/motion';
import { GridIndex } from './types';

export default class RopePhysicsModeller {
  public static calculateVisitedPositions(
    motions: Motion[],
    knotCount: number
  ): number {
    let knots: GridIndex[] = this.initKnots(knotCount);
    const tailVisitedIndexes = this.calculateTailVisitedIndexes(motions, knots);
    return this.getUniqueLength(tailVisitedIndexes);
  }

  private static initKnots(knotCount: number) {
    return [...new Array(knotCount)].map((value) => {
      return { row: 0, col: 0 };
    });
  }

  private static calculateTailVisitedIndexes(
    motions: Motion[],
    knots: GridIndex[]
  ) {
    return motions
      .map((motion: Motion) => {
        const movementData = motion.execute(knots);
        knots = [...movementData.knotPositions];
        return [...movementData.tailVisitedPositions];
      })
      .reduce(RopePhysicsModeller.aggregateVisitedIndexesReducer, []);
  }

  private static aggregateVisitedIndexesReducer(
    acc: GridIndex[],
    visitedPositions: GridIndex[]
  ): GridIndex[] {
    return [...acc, ...visitedPositions];
  }

  private static getUniqueLength(tailVisitedIndexes: GridIndex[]) {
    return new Set(
      tailVisitedIndexes.map((gridIndex) => `${gridIndex.row},${gridIndex.col}`)
    ).size;
  }
}
