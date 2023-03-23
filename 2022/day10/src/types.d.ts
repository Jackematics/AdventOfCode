export type ScreenDims = {
  width: number;
  height: number;
};

export type InstructionOutput = {
  X: number;
  cycle: number;
  signalStrength: number;
  cycleIntervals: number[];
  screenDims: ScreenDims;
  crtOutput: string[];
};
