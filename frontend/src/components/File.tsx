import React from "react";
import { cn } from "@/lib/utils";
import { backend } from "wailsjs/go/models";
import { FILE_HEIGHT } from "@/screens/FileBrowser";
import { createUrl } from "@/lib/url";
const FileItem = ({ className, file, style }: Props) => {
	if (!file) return null;
	const url = createUrl(file.url, { size: FILE_HEIGHT });
	console.log({ url });
	return (
		<div
			className={cn("flex-1 justify-center items-center", className)}
			style={{ height: FILE_HEIGHT, ...style }}
		>
			<img
				src={url}
				alt={file.name}
				loading="lazy"
				className="flex m-auto  rounded-md overflow-hidden shadow-sm aspect-auto object-contain"
				style={{ height: FILE_HEIGHT }}
			/>
		</div>
	);
};

export default FileItem;

type Props = {
	className?: string;
	file: backend.File;
	style?: React.CSSProperties;
};
