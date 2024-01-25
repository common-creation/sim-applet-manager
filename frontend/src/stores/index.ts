import { createContext, useContext } from "react";


import { configure } from "mobx";
import { SimStore } from './SimStore';
import { LogStore } from './LogStore';

configure({
  enforceActions: "never",
});

const store = {
  SimStore: SimStore(),
  LogStore: LogStore,
};

store.LogStore.watch();

export const StoreContext = createContext(store);

export const useStore = () => {
  return useContext<typeof store>(StoreContext);
};

export default store;