import { createContext } from "react";
import store from "../index";

// react-hook 与 mobx下结合的分配形式，原先的inject方式只能用于class，不适用于function.
// export const StoreContext = createContext<ActList>({} as ActList);
export const StoreContext = createContext(store);
export const StoreProvider = StoreContext.Provider;
