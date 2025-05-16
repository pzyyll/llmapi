export const { format: formatNumber } = Intl.NumberFormat('en-GB', {
	notation: 'compact',
	maximumFractionDigits: 1
});

export const { format: formatDate } = Intl.DateTimeFormat('zh-CN', {
	year: 'numeric',
	month: '2-digit',
	day: '2-digit',
	hour: '2-digit',
	minute: '2-digit',
	second: '2-digit'
});

export function formatDateTimeString(dateString: string): string {
	const date = new Date(dateString);
	// convert to local date
	const options: Intl.DateTimeFormatOptions = {
		year: 'numeric',
		month: '2-digit',
		day: '2-digit',
		hour: '2-digit',
		minute: '2-digit',
		second: '2-digit'
	};
	const formatter = new Intl.DateTimeFormat('zh', options);
	const parts = formatter.formatToParts(date);
	const formattedDate = parts
		.map((part) => {
			if (part.type === 'literal') {
				return part.value;
			}
			return part.value.padStart(2, '0');
		})
		.join('');
	return formattedDate.replace(/(\d{4})(\d{2})(\d{2})(\d{2})(\d{2})(\d{2})/, '$1-$2-$3 $4:$5:$6');
}
