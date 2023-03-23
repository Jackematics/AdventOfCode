import CathodRayTube from './cathod-ray-tube';
import { IInstruction, Addx, Noop } from './instruction';
import FileProcessor from './file-processor';
import { describe, it, expect } from '@jest/globals';

describe('CathodeRayTube', () => {
  it('should return the aggregate signal strength after processing instructions with one interval', () => {
    const instructions: IInstruction[] = [
      new Addx(15),
      new Addx(-11),
      new Addx(6),
      new Addx(-3),
      new Addx(5),
      new Addx(-1),
      new Addx(-8),
      new Addx(13),
      new Addx(4),
      new Noop(),
      new Addx(-1),
      new Addx(5),
      new Addx(-1),
      new Addx(5),
      new Addx(-1),
      new Addx(5),
      new Addx(-1),
      new Addx(5),
      new Addx(-1),
      new Addx(-35),
    ];

    const cycleIntervals = [20];

    expect(
      CathodRayTube.calculateAggregatedOutput(instructions, cycleIntervals, {
        width: 40,
        height: 6,
      }).signalStrength
    ).toBe(420);
  });

  it('should return the aggregate signal strength after processing instructions with multiple intervals', () => {
    const instructions: IInstruction[] =
      FileProcessor.processInputIntoInstructions('src/example-input.txt');

    const cycleIntervals = [20, 60, 100, 140, 180, 220];

    expect(
      CathodRayTube.calculateAggregatedOutput(instructions, cycleIntervals, {
        width: 40,
        height: 6,
      }).signalStrength
    ).toBe(13140);
  });

  it('should count draw the screen correctly ', () => {
    const instructions: IInstruction[] =
      FileProcessor.processInputIntoInstructions('src/example-input.txt');

    const screenDims = { width: 40, height: 6 };
    const expected = [
      '##..##..##..##..##..##..##..##..##..##..',
      '###...###...###...###...###...###...###.',
      '####....####....####....####....####....',
      '#####.....#####.....#####.....#####.....',
      '######......######......######......####',
      '#######.......#######.......#######.....',
    ];

    expect(
      CathodRayTube.calculateAggregatedOutput(instructions, [], screenDims)
        .crtOutput
    ).toStrictEqual(expected);
  });
});
