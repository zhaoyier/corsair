import * as React from "react";
import ParcelListStore from "../store/parcelList";
import ParcelDetailStore from "../store/parcelDetail";
import ModaltStore from "../store/modal";

/**
 * mobx-react-lite + context 使用
 * 1. 创建 store --- mobx
 * 2. 创建 Context
 * 创建 StoresContext 并导出，那么后代组件在 useContext 时便可以得到包含 ParcelListStore, ParcelDetailStore 实例的对象
 * 3. 由于多组件都需要使用 useContext，我们将其封装为 hook 函数  useStores
 * 通过 useStores 获取 React.createContext 给的初始值对象
 * （前提是没有 StoresContext.Provider 组件，如果使用了该组件，则拿到的是 StoresContext.Provider 的 value 属性对应的值）
 * 4.后代组件获取 store 并调用 action 修改状态，使用 mobx-react-lite 更新组件
 */
//

export const parcelStore = {
	parcelListStore: new ParcelListStore(),
	parcelDetailStore: new ParcelDetailStore(),
	modalStore: new ModaltStore()
};

export const StoresContext = React.createContext(parcelStore);
export const useStores = () => React.useContext(StoresContext);

/**
 * 1.使用封装好的hooks  useStores
 *
 * import React from 'react';
 * import { useObserver } from 'mobx-react-lite';
 *
 * const Counter = () => {
 *   let store = useStores(); // 获取store
 *
 *   const {counterStore, themeStore} = store;
 *
 *   const handleIncrement = () => {
 *     counterStore.increment();
 *   *  }
 *   const handleDecrement = () => {
 *     counterStore.decrement();
 *  }
 *
 *   return useObserver(() => (
 *    <div>
 *       <p>count: {counterStore.count}</p>
 *       <p>theme: {themeStore.theme}</p>
 *       <button onClick={handleIncrement}>add</button>
 *       <button onClick={handleDecrement}>dec</button>
 *     </div>
 *   ))
 * }
 *
 * export default Counter;
 *
 *
 * 2. 使用 StoresContext.Provider组件
 * export const StoreProvider = props => {
 * 	return <StoresContext.Provider value={parcelStore}>{props.children}</StoresContext.Provider>;
 * };
 */
