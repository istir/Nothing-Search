import React, { useRef } from "react";
// import { cn } from "@/lib/utils";
import { useQuery } from "@tanstack/react-query";
import { cn } from "@/lib/utils";
import { LoadFiles } from "!/App";
import File from "@/components/File";
import { useVirtualizer } from "@tanstack/react-virtual";

export const FILE_HEIGHT = 200;
const FileBrowser = ({ className }: Props) => {
	const scrollRef = useRef<HTMLDivElement>(null);
	const files = useQuery({
		queryKey: ["files"],
		queryFn: () =>
			LoadFiles(
				"/Users/lina/Library/Mobile Documents/com~apple~CloudDocs/Lina/hot",
			),
	});

	const fileVirtualizer = useVirtualizer({
		count: files.data?.length || 0,
		estimateSize: () => FILE_HEIGHT,
		overscan: 25,
		getScrollElement: () => scrollRef.current,
	});

	console.log("files", files.data?.length);
	console.log("virtual items", fileVirtualizer.getVirtualItems());

	return (
		<>
			<div
				ref={scrollRef}
				className="List h-dvh"
				style={{
					// height: `200px`,
					// width: `400px`,
					overflow: "auto",
				}}
			>
				<div
					style={{
						height: `${fileVirtualizer.getTotalSize()}px`,
						width: "100%",
						position: "relative",
					}}
				>
					{fileVirtualizer.getVirtualItems().map((virtualRow) =>
						files.data?.[virtualRow.index] ? (
							<File
								file={files.data?.[virtualRow.index]}
								key={virtualRow.index}
								style={{
									position: "absolute",
									top: 0,
									left: 0,
									width: "100%",
									height: `${virtualRow.size}px`,
									transform: `translateY(${virtualRow.start}px)`,
								}}
							/>
						) : null,
					)}
				</div>
			</div>
		</>
	);

	// return (
	// 	<div
	// 		ref={scrollRef}
	// 		style={{ overflowY: "auto", border: "1px solid #c8c8c8" }}
	// 	>
	// 		<div
	// 			className={cn(" gap-5 grid p-5", className)}
	// 			// style={{
	// 			// 	gridTemplateColumns: "repeat(auto-fit, minmax(240px, 1fr))",
	// 			// }}
	// 			style={{
	// 				height: `${fileVirtualizer.getTotalSize()}px`,
	// 				position: "relative",
	// 				// gridTemplateColumns: "repeat(auto-fit, minmax(240px, 1fr))",
	// 				// display: "grid",
	// 			}}
	// 		>
	// 			{fileVirtualizer
	// 				.getVirtualItems()
	// 				.map((virtualItem) =>
	// 					files.data?.[virtualItem.index] ? (
	// 						<File
	// 							file={files.data?.[virtualItem.index]}
	// 							key={virtualItem.index}
	// 						/>
	// 					) : null,
	// 				)}
	// 			{/* {files.data?.map((file) => (
	// 				<File file={file} key={file.url} />
	// 			))} */}
	// 		</div>
	// 	</div>
	// );
};

export default FileBrowser;

type Props = {
	className?: string;
};
