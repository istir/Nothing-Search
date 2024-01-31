import React from "react";
import { cn } from "@/lib/utils";
const File = ({ className }: Props) => {
  return <div className={cn("flex-1", className)}>ball</div>;
};

export default File;

type Props = {
  className?: string;
};
