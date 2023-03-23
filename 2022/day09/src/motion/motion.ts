import { Direction, TailMoveDirection as KnotMoveDirection } from '../enum';
import { GridIndex, MotionData } from '../types';

export default class Motion {
  public readonly direction: Direction;
  public readonly distance: number;

  constructor(direction: Direction, distance: number) {
    this.direction = direction;
    this.distance = distance;
  }

  public execute(knots: GridIndex[]): MotionData {
    const headVisitedIndexes = this.calculateVisitedHeadIndexes(knots[0]);

    let currentKnotIndex = 1;
    const allKnotsVisitedIndexes = this.recursivelyChartKnots(
      [headVisitedIndexes],
      knots,
      currentKnotIndex
    );

    return {
      knotPositions: allKnotsVisitedIndexes.map(
        (visitedPositions) => visitedPositions[visitedPositions.length - 1]
      ),
      tailVisitedPositions:
        allKnotsVisitedIndexes[allKnotsVisitedIndexes.length - 1],
    };
  }

  private calculateVisitedHeadIndexes(
    initialHeadIndex: GridIndex
  ): GridIndex[] {
    return Array.from(Array(this.distance + 1).keys()).map((i) => {
      return {
        [Direction.Up]: {
          row: initialHeadIndex.row + i,
          col: initialHeadIndex.col,
        },
        [Direction.Right]: {
          row: initialHeadIndex.row,
          col: initialHeadIndex.col + i,
        },
        [Direction.Down]: {
          row: initialHeadIndex.row - i,
          col: initialHeadIndex.col,
        },
        [Direction.Left]: {
          row: initialHeadIndex.row,
          col: initialHeadIndex.col - i,
        },
      }[this.direction];
    });
  }

  private recursivelyChartKnots(
    allKnotsVisitedIndexes: GridIndex[][],
    knots: GridIndex[],
    currentKnotIndex: number
  ) {
    allKnotsVisitedIndexes.push(
      this.addNextKnotVisitedIndexes(
        allKnotsVisitedIndexes,
        knots,
        currentKnotIndex
      )
    );

    if (allKnotsVisitedIndexes.length < knots.length)
      this.recursivelyChartKnots(
        allKnotsVisitedIndexes,
        knots,
        currentKnotIndex + 1
      );

    return allKnotsVisitedIndexes;
  }

  private addNextKnotVisitedIndexes(
    allKnotsVisitedIndexes: GridIndex[][],
    knots: GridIndex[],
    currentKnotIndex: number
  ) {
    const prevKnotPositions =
      allKnotsVisitedIndexes[allKnotsVisitedIndexes.length - 1];
    return this.calculateNextKnotVisitedIndexes(
      knots,
      currentKnotIndex,
      prevKnotPositions
    );
  }

  private calculateNextKnotVisitedIndexes(
    knots: GridIndex[],
    currentKnotIndex: number,
    prevKnotPositions: GridIndex[]
  ): GridIndex[] {
    let currentKnot = {
      ...knots[currentKnotIndex],
    };
    return prevKnotPositions.map((prevKnot) => {
      currentKnot = this.calculateVisitedIndex(prevKnot, currentKnot);
      return this.calculateVisitedIndex(prevKnot, currentKnot);
    });
  }

  private calculateVisitedIndex(prevKnot: GridIndex, currentKnot: GridIndex) {
    const tailMoveDirection = this.calculateKnotMoveDirection(
      prevKnot,
      currentKnot
    );
    return this.calculateNextTailPosition(tailMoveDirection, currentKnot);
  }

  private calculateKnotMoveDirection(
    prevKnotIndex: GridIndex,
    currentKnotIndex: GridIndex
  ): KnotMoveDirection {
    const rowColDiffKey = `${prevKnotIndex.row - currentKnotIndex.row},${
      prevKnotIndex.col - currentKnotIndex.col
    }`;

    return (
      {
        ['1,-2']: KnotMoveDirection.UpLeft,
        ['2,-2']: KnotMoveDirection.UpLeft,
        ['2,-1']: KnotMoveDirection.UpLeft,
        ['2,0']: KnotMoveDirection.Up,
        ['2,1']: KnotMoveDirection.UpRight,
        ['2,2']: KnotMoveDirection.UpRight,
        ['1,2']: KnotMoveDirection.UpRight,
        ['0,2']: KnotMoveDirection.Right,
        ['-1,2']: KnotMoveDirection.DownRight,
        ['-2,2']: KnotMoveDirection.DownRight,
        ['-2,1']: KnotMoveDirection.DownRight,
        ['-2,0']: KnotMoveDirection.Down,
        ['-2,-1']: KnotMoveDirection.DownLeft,
        ['-2,-2']: KnotMoveDirection.DownLeft,
        ['-1,-2']: KnotMoveDirection.DownLeft,
        ['0,-2']: KnotMoveDirection.Left,
      }[rowColDiffKey] ?? KnotMoveDirection.None
    );
  }

  private calculateNextTailPosition(
    tailMoveDirection: KnotMoveDirection,
    currentTailPosition: GridIndex
  ): GridIndex {
    return {
      [KnotMoveDirection.None]: currentTailPosition,
      [KnotMoveDirection.UpLeft]: {
        row: currentTailPosition.row + 1,
        col: currentTailPosition.col - 1,
      },
      [KnotMoveDirection.Up]: {
        row: currentTailPosition.row + 1,
        col: currentTailPosition.col,
      },
      [KnotMoveDirection.UpRight]: {
        row: currentTailPosition.row + 1,
        col: currentTailPosition.col + 1,
      },
      [KnotMoveDirection.Right]: {
        row: currentTailPosition.row,
        col: currentTailPosition.col + 1,
      },
      [KnotMoveDirection.DownRight]: {
        row: currentTailPosition.row - 1,
        col: currentTailPosition.col + 1,
      },
      [KnotMoveDirection.Down]: {
        row: currentTailPosition.row - 1,
        col: currentTailPosition.col,
      },
      [KnotMoveDirection.DownLeft]: {
        row: currentTailPosition.row - 1,
        col: currentTailPosition.col - 1,
      },
      [KnotMoveDirection.Left]: {
        row: currentTailPosition.row,
        col: currentTailPosition.col - 1,
      },
    }[tailMoveDirection];
  }
}
