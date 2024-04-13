import FileItem from "@/components/File";
import type { backend } from "wailsjs/go/models";

export function FileBrowserCell({
	columnIndex,
	rowIndex,
	style,
	data,
}: FileBrowserCellProps) {
	return (
		<div style={style}>
			<FileItem file={data} />
		</div>
	);
}
type FileBrowserCellProps = {
	columnIndex: number;
	rowIndex: number;
	style: React.CSSProperties;
	data: backend.File;
};
