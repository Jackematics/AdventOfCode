import CalorieCounter from './calorie-counter';

describe('CalorieCounter()', () => {
  it('should return the max value of an array of elfInventories', () => {
    const elfInventories = [
      [1000, 2000, 3000],
      [4000],
      [5000, 6000],
      [7000, 8000, 9000],
      [10000],
    ];

    expect(CalorieCounter.calculateMaxInventoryValue(elfInventories)).toBe(
      24000
    );
  });

  it('should return the sum of the top three elfInventories', () => {
    const elfInventories = [
      [1000, 2000, 3000],
      [4000],
      [5000, 6000],
      [7000, 8000, 9000],
      [10000],
    ];

    expect(CalorieCounter.calculateSumOfTopThree(elfInventories)).toBe(45000);
  });
});
