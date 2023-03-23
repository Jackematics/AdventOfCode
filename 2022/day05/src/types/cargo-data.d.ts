import Cargo from '../cargo';
import MoveOrder from './move-order';

type CargoData = {
  cargo: Cargo;
  moveOrders: MoveOrder[];
};

export default CargoData;
