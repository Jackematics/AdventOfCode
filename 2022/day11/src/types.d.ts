export type MonkeyState = {
  items: number[];
  divisor: number;
  operation: (oldWorry: number) => number;
  test: (worry: number) => number;
};

export type ThrownItem = {
  monkey: number;
  item: number;
};
