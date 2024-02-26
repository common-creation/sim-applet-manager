import { makeAutoObservable } from "mobx";
import ja from "../assets/i18n/ja.json";
import en from "../assets/i18n/en.json";

export const I18nStore = () => {
  return makeAutoObservable({
    t(key: keyof typeof ja, ...args: string[]) {
      let base;
      if (navigator.language == "ja") {
        base = ja[key];
      } else {
        base = en[key];
      }
      if (!base) {
        base = key;
      }

      if (args && args.length > 0) {
        for (let i = 0; i < args.length; i++) {
          base = base.replace(`{${i}}`, args[i]);
        }
      }

      return base;
    },
  });
};