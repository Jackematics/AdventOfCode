export type GridIndex = {
  row: number;
  col: number;
};

export type MotionData = {
  knotPositions: GridIndex[];
  tailVisitedPositions: GridIndex[];
};
