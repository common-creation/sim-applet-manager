import { createContext, useContext } from "react";


import { configure } from "mobx";
import { SimStore } from "./SimStore";
import { LogStore } from "./LogStore";
import { I18nStore } from "./I18nStore";

configure({
  enforceActions: "never",
});

const store = {
  SimStore: SimStore(),
  LogStore: LogStore,
  I18nStore: I18nStore(),
};

store.LogStore.watch();

export const StoreContext = createContext(store);

export const useStore = () => {
  return useContext<typeof store>(StoreContext);
};

export default store;