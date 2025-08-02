<script>
	let { domain, hasLogo } = $props();
	
	const logoUrls = {
		'privacy-guide.org': '/placeholder.svg?height=32&width=32',
		'digital-rights.org': '/placeholder.svg?height=32&width=32',
		'secure-chat.info': '/placeholder.svg?height=32&width=32'
	};
	
	function getLogoUrl(domain) {
		return logoUrls[domain] || null;
	}
	
	function handleImageError(event) {
		event.target.style.display = 'none';
		event.target.nextElementSibling.style.display = 'flex';
	}
</script>

<div class="flex-shrink-0">
	{#if hasLogo && getLogoUrl(domain)}
		<div class="relative h-8 w-8">
			<img
				src={getLogoUrl(domain) || "/placeholder.svg"}
				alt={`${domain} logo`}
				class="h-8 w-8 rounded-lg object-cover"
				onerror={handleImageError}
			/>
			<div class="hidden h-8 w-8 items-center justify-center rounded-lg bg-gradient-to-br from-purple-500 to-purple-700 text-white font-semibold text-sm">
				B
			</div>
		</div>
	{:else}
		<div class="h-8 w-8 flex items-center justify-center rounded-lg bg-gradient-to-br from-purple-500 to-purple-700 text-white font-semibold text-sm">
			B
		</div>
	{/if}
</div>
