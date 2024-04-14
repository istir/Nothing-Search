import { LoadAllFilesToDb } from "!/App";
import FileBrowserItems from "@/components/FileBrowserItems";
import { useEffect } from "react";

export function FileBrowser() {
  // useEffect(()=>{
  //   LoadAllFilesToDb()
  // },[])
  return <FileBrowserItems />;
}
