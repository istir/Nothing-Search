import React from "react";
import { cn } from "@/lib/utils";
import { useQuery } from "@tanstack/react-query";
import { LoadFiles } from "wailsjs/go/main/App";

const FileBrowser = ({ className }: Props) => {
	const files = useQuery({
		queryKey: ["files"],
		queryFn: () =>
			LoadFiles(
				"/Users/lina/Library/Mobile Documents/com~apple~CloudDocs/Lina/memes/cool",
			),
	});
	return <div className={cn("flex-1", className)}>cock</div>;
};

export default FileBrowser;

type Props = {
	className?: string;
};
