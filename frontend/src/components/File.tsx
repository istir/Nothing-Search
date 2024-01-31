import React from "react";
import { cn } from "@/lib/utils";
import { backend } from "wailsjs/go/models";
import { FILE_HEIGHT } from "@/screens/FileBrowser";
const File = ({ className, file, style }: Props) => {
	if (!file) return null;
	return (
		<div
			className={cn("flex-1 justify-center items-center", className)}
			style={{ height: FILE_HEIGHT, ...style }}
		>
			<img
				src={file.url}
				alt={file.name}
				loading="lazy"
				className="flex m-auto  rounded-md overflow-hidden shadow-sm"
				style={{ height: FILE_HEIGHT }}
			/>
		</div>
	);
};

export default File;

type Props = {
	className?: string;
	file: backend.File;
	style?: React.CSSProperties;
};
