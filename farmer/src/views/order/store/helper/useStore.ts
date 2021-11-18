import { useContext } from "react";
import { StoreContext } from "./storeProvider";

// 从store的provider中拿出来用react的useContext来注入。
// export const ActivityListStore = () => useContext(StoreContext);
// tslint:disable-next-line: react-hooks-nesting
export const store = () => useContext(StoreContext);
