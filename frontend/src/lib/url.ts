export function createUrl(
	baseUrl: string,
	queryData: Record<string, string | number>,
): string {
	// const url = new URL(baseUrl);
	let url = baseUrl;
	let separator = "?";
	if (url.includes("?")) {
		separator = "&";
	}

	for (const [key, value] of Object.entries(queryData)) {
		url += `${separator}${key}=${value.toString()}`;
		separator = "&";
	}
	return url.toString();
}
