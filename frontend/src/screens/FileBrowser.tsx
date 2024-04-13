import React, { useRef, useState } from "react";
import { FixedSizeGrid as Grid } from "react-window";

// import { cn } from "@/lib/utils";
import { useQuery } from "@tanstack/react-query";
import { cn } from "@/lib/utils";
import { LoadFiles } from "!/App";
import File from "@/components/File";
import { FileBrowserCell } from "@/components/FileBrowserList";
import ReactVirtualizedAutoSizer from "react-virtualized-auto-sizer";

export const FILE_HEIGHT = 200;
const FileBrowser = ({ className }: Props) => {
	const scrollRef = useRef<HTMLDivElement>(null);
	const files = useQuery({
		queryKey: ["files"],
		queryFn: () =>
			LoadFiles(
				"/Users/lina/Library/Mobile Documents/com~apple~CloudDocs/Lina/pretty",
			),
	});

	if (!files.data?.length) {
		return <div>No data</div>;
	}
	return (
		<div
			style={{
				height: "100vh",
				width: "100vw",
			}}
		>
			<ReactVirtualizedAutoSizer>
				{({ height, width }) => (
					<Grid
						columnCount={6}
						columnWidth={width / 6}
						height={height}
						rowCount={(files.data?.length || 6) / 6}
						rowHeight={FILE_HEIGHT}
						width={width}
					>
						{({ columnIndex, rowIndex, style, isScrolling }) => (
							<FileBrowserCell
								columnIndex={columnIndex}
								rowIndex={rowIndex}
								style={style}
								data={files.data?.[columnIndex * rowIndex]}
							/>
						)}
					</Grid>
				)}
			</ReactVirtualizedAutoSizer>
			)
		</div>
	);
};

export default FileBrowser;

type Props = {
	className?: string;
};
