import AssignmentPair from './assignment-pair';
import CampCleanupOrganiser from './camp-cleaner-organiser';
import TestData from './test-data';

describe('CampCleanupOrganiser()', () => {
  it.each([
    {
      input: [TestData.assignmentPairs[0]],
      expected: 0,
    },
    {
      input: [TestData.assignmentPairs[1]],
      expected: 0,
    },
    {
      input: [TestData.assignmentPairs[2]],
      expected: 0,
    },
    {
      input: [TestData.assignmentPairs[3]],
      expected: 1,
    },
    {
      input: [TestData.assignmentPairs[4]],
      expected: 1,
    },
    {
      input: [TestData.assignmentPairs[5]],
      expected: 0,
    },
  ])(
    `should return the return the count of a pair that have completely duplicated sections`,
    ({ input, expected }) => {
      expect(
        CampCleanupOrganiser.calculateTotalFullSectionDuplication(input)
      ).toBe(expected);
    }
  );

  it('should return the return the count of multiple pairs that have completely duplicated sections', () => {
    expect(
      CampCleanupOrganiser.calculateTotalFullSectionDuplication(
        TestData.assignmentPairs
      )
    ).toBe(2);
  });

  it.each([
    {
      input: [TestData.assignmentPairs[0]],
      expected: 0,
    },
    {
      input: [TestData.assignmentPairs[1]],
      expected: 0,
    },
    {
      input: [TestData.assignmentPairs[2]],
      expected: 1,
    },
    {
      input: [TestData.assignmentPairs[3]],
      expected: 1,
    },
    {
      input: [TestData.assignmentPairs[4]],
      expected: 1,
    },
    {
      input: [TestData.assignmentPairs[5]],
      expected: 1,
    },
  ])(
    `should return the return the count of a pair that have partially duplicated sections`,
    ({ input, expected }) => {
      expect(
        CampCleanupOrganiser.calculateTotalPartialSectionDuplication(input)
      ).toBe(expected);
    }
  );

  it('should return the return the count of multiple pairs that have partially duplicated sections', () => {
    expect(
      CampCleanupOrganiser.calculateTotalPartialSectionDuplication(
        TestData.assignmentPairs
      )
    ).toBe(4);
  });
});
