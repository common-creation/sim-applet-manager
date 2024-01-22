import { createContext, useContext } from "react";


import { configure } from "mobx";
import { SimStore } from './SimStore';

configure({
  enforceActions: "never",
});

const store = {
  SimStore: SimStore(),
};

export const StoreContext = createContext(store);

export const useStore = () => {
  return useContext<typeof store>(StoreContext);
};

export default store;