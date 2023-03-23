import { Direction } from './enum';
import FileProcessor from './file-processor';
import Motion from './motion/motion';

describe('FileProcessor()', () => {
  it('should convert a text file of lines into an array of motions', () => {
    const data: Motion[] = FileProcessor.processInput('src/example-input.txt');
    const expected = [
      {
        direction: Direction.Right,
        distance: 4,
      },
      {
        direction: Direction.Up,
        distance: 4,
      },
      {
        direction: Direction.Left,
        distance: 3,
      },
      {
        direction: Direction.Down,
        distance: 1,
      },
      {
        direction: Direction.Right,
        distance: 4,
      },
      {
        direction: Direction.Down,
        distance: 1,
      },
      {
        direction: Direction.Left,
        distance: 5,
      },
      {
        direction: Direction.Right,
        distance: 2,
      },
    ];

    data.forEach((motion, index) => {
      expect(motion.direction).toBe(expected[index].direction);
      expect(motion.distance).toBe(expected[index].distance);
    });
  });
});
