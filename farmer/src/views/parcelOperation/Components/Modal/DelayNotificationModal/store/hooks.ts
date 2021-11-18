import DelayNotificationModalStore from "./store";
import * as React from "react";

export const store = new DelayNotificationModalStore();

export const StoresContext = React.createContext(store);
export const useModalStores = () => React.useContext(StoresContext);
