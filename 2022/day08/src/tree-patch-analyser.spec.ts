import TreePatchAnalyser from './tree-patch-analyser';

describe('TreePatchAnalyser()', () => {
  it('should return the number of trees visible in a patch', () => {
    const treePatch = [
      [3, 0, 3, 7, 3],
      [2, 5, 5, 1, 2],
      [6, 5, 3, 3, 2],
      [3, 3, 5, 4, 9],
      [3, 5, 3, 9, 0],
    ];

    expect(TreePatchAnalyser.visibleTreesCount(treePatch)).toBe(21);
  });

  it('should return the highest scenic score in a patch', () => {
    const treePatch = [
      [3, 0, 3, 7, 3],
      [2, 5, 5, 1, 2],
      [6, 5, 3, 3, 2],
      [3, 3, 5, 4, 9],
      [3, 5, 3, 9, 0],
    ];

    expect(TreePatchAnalyser.calculateMaxScenicScore(treePatch)).toBe(8);
  });
});
