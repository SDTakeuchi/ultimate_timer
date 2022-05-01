function zeroPadding(num: number, len: number): string{
	return ( Array(len).join('0') + num ).slice( -len );
}

export default zeroPadding;