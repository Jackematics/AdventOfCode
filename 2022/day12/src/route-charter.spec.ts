import FileProcessor from './file-processor';

describe('RouteCharter', () => {
  it('should return the closest route to the objective', () => {
    const heightMap = FileProcessor.processInput('src/example-input.txt');

    const route = RouteCharter.chartRouteToObjective(heightMap);

    expect(route.length).toBe(31);
  });
});
