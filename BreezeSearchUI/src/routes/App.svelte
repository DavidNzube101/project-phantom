<script>
	import { Search, Shield, Zap, Eye, EyeOff, Lock, Sparkles, TrendingUp, Clock, ChevronLeft, ChevronRight, Download, FileText, FileSpreadsheet, X, Bot, ChevronDown, Mic, MicOff } from 'lucide-svelte';
	import { Button } from '$lib/components/ui/button';
	import { Input } from '$lib/components/ui/input';
	import { Card, CardContent } from '$lib/components/ui/card';
	import { Badge } from '$lib/components/ui/badge';
	import SiteLogo from './components/SiteLogo.svelte';
	import AISummary from './components/AISummary.svelte';
	import NewsCarousel from './components/NewsCarousel.svelte';
	import CurrencyExchange from './components/CurrencyExchange.svelte';
	import WeatherWidget from './components/WeatherWidget.svelte';
	import SpecificAnswer from './components/SpecificAnswer.svelte';
	import SolanaExplorer from './components/SolanaExplorer.svelte';
	import TokenPriceChart from './components/TokenPriceChart.svelte';
	import SearchResultSkeleton from './components/SearchResultSkeleton.svelte'; // New import

	let searchQuery = $state('');
	let isSearching = $state(false);
	let hasSearched = $state(false);
	let showPrivacyFeatures = $state(false);
	let currentPage = $state(1);
	let showSuggestions = $state(false);
	let suggestions = $state([]);
	let showExportMenu = $state(false);
	let isListening = $state(false);
	let recognition = $state(null);

	// Conditional widget states
	let showCurrencyExchange = $state(false);
	let showWeather = $state(false);
	let showSpecificAnswer = $state(false);
	let specificAnswerData = $state(null);
	let showSolanaExplorer = $state(false);
	let solanaQueryType = $state(''); // 'address' or 'transaction'
	let solanaQuery = $state('');
	let showTokenPriceChart = $state(false);
	let tokenChartSymbol = $state('');

	let resultsLoading = $state(false); // New state for skeleton loading

	const trendingSearches = [
		'Secure messaging apps',
		'Privacy tools 2024',
		'Open source software',
		'Digital privacy rights',
		'Encrypted cloud storage'
	];

	const recentSearches = [
		'VPN comparison guide',
		'Privacy-focused browsers',
		'Data protection laws'
	];

	const allSuggestions = [
		'secure messaging apps',
		'secure email providers',
		'security best practices',
		'privacy tools 2024',
		'privacy-focused browsers',
		'privacy laws',
		'open source software',
		'open source alternatives',
		'digital privacy rights',
		'data protection',
		'encrypted cloud storage',
		'encryption tools',
		'USD to EUR',
		'weather in New York',
		'what is bitcoin',
		'how to encrypt files',
		'SOL price',
		'BTC chart',
		'ETH price'
	];

	const mockResults = [
		{
			title: 'Complete Guide to Digital Privacy in 2024',
			url: 'privacy-guide.org',
			domain: 'privacy-guide.org',
			description: 'Learn essential privacy practices, tools, and techniques to protect your digital footprint. Comprehensive coverage of encryption, secure communications, and data protection.',
			timestamp: '2h ago',
			verified: true,
			hasLogo: true
		},
		{
			title: 'Open Source Privacy Tools You Should Know',
			url: 'opensource-privacy.com',
			domain: 'opensource-privacy.com',
			description: 'Discover powerful open-source alternatives to mainstream apps that respect your privacy. From browsers to messaging apps and beyond.',
			timestamp: '4h ago',
			verified: true,
			hasLogo: false
		},
		{
			title: 'Understanding Data Collection and Your Rights',
			url: 'digital-rights.org',
			domain: 'digital-rights.org',
			description: 'A deep dive into how companies collect your data, what rights you have, and practical steps to limit unwanted data harvesting.',
			timestamp: '1d ago',
			verified: false,
			hasLogo: true
		},
		{
			title: 'Best VPN Services for Privacy in 2024',
			url: 'vpn-reviews.net',
			domain: 'vpn-reviews.net',
			description: 'Comprehensive review of the most trusted VPN services that prioritize user privacy and security. Compare features, pricing, and performance.',
			timestamp: '3h ago',
			verified: true,
			hasLogo: false
		},
		{
			title: 'Encrypted Messaging Apps Comparison',
			url: 'secure-chat.info',
			domain: 'secure-chat.info',
			description: 'Compare the security features of popular encrypted messaging apps. Learn about end-to-end encryption, metadata protection, and more.',
			timestamp: '5h ago',
			verified: true,
			hasLogo: true
		}
	];

	const totalPages = 12;
	const resultsPerPage = 10;

	// Initialize speech recognition
	function initSpeechRecognition() {
		if ('webkitSpeechRecognition' in window || 'SpeechRecognition' in window) {
			const SpeechRecognition = window.SpeechRecognition || window.webkitSpeechRecognition;
			recognition = new SpeechRecognition();
			recognition.continuous = false;
			recognition.interimResults = false;
			recognition.lang = 'en-US';

			recognition.onstart = () => {
				isListening = true;
			};

			recognition.onresult = (event) => {
				const transcript = event.results[0][0].transcript;
				searchQuery = transcript;
				isListening = false;
				handleSearch();
			};

			recognition.onerror = () => {
				isListening = false;
			};

			recognition.onend = () => {
				isListening = false;
			};
		}
	}

	function startVoiceSearch() {
		if (recognition) {
			recognition.start();
		} else {
			initSpeechRecognition();
			if (recognition) {
				recognition.start();
			}
		}
	}

	function stopVoiceSearch() {
		if (recognition) {
			recognition.stop();
		}
		isListening = false;
	}

	function detectSearchType(query) {
		const trimmedQuery = query.trim();
		const lowerQuery = trimmedQuery.toLowerCase();
		
		// Reset all conditions
		showCurrencyExchange = false;
		showWeather = false;
		showSpecificAnswer = false;
		specificAnswerData = null;
		showSolanaExplorer = false;
		solanaQueryType = '';
		solanaQuery = '';
		showTokenPriceChart = false;
		tokenChartSymbol = '';

		// Solana blockchain detection (Base58 is case-sensitive, so use original query)
		const solanaAddressPattern = /^[1-9A-HJ-NP-Za-km-z]{32,44}$/;
		const solanaTransactionPattern = /^[1-9A-HJ-NP-Za-km-z]{80,90}$/; // Typical length for Solana tx hash is 88 chars

		// Check for transaction first as it's longer and more specific
		if (solanaTransactionPattern.test(trimmedQuery)) {
			showSolanaExplorer = true;
			solanaQueryType = 'transaction';
			solanaQuery = trimmedQuery;
		} else if (solanaAddressPattern.test(trimmedQuery)) {
			showSolanaExplorer = true;
			solanaQueryType = 'address';
			solanaQuery = trimmedQuery;
		}

		// Currency exchange detection
		const currencyKeywords = ['usd', 'eur', 'gbp', 'jpy', 'btc', 'eth', 'sol', 'exchange rate', 'currency', 'convert', 'to'];
		const currencyPattern = /(\w{3})\s+to\s+(\w{3})|exchange rate|currency|convert/i;
		if (currencyKeywords.some(keyword => lowerQuery.includes(keyword)) || currencyPattern.test(lowerQuery)) {
			showCurrencyExchange = true;
		}

		// Weather detection
		const weatherKeywords = ['weather', 'temperature', 'forecast', 'rain', 'sunny', 'cloudy', 'snow', 'climate'];
		if (weatherKeywords.some(keyword => lowerQuery.includes(keyword))) {
			showWeather = true;
		}

		// Token Price Chart detection
		const tokenChartKeywords = [
			{ keyword: 'sol price', symbol: 'SOL' },
			{ keyword: 'sol chart', symbol: 'SOL' },
			{ keyword: 'btc price', symbol: 'BTC' },
			{ keyword: 'btc chart', symbol: 'BTC' },
			{ keyword: 'eth price', symbol: 'ETH' },
			{ keyword: 'eth chart', symbol: 'ETH' }
		];
		for (const { keyword, symbol } of tokenChartKeywords) {
			if (lowerQuery.includes(keyword)) {
				showTokenPriceChart = true;
				tokenChartSymbol = symbol;
				break;
			}
		}

		// Specific answer detection
		const questionPatterns = [
			{ pattern: /what is (.*)/i, type: 'definition' },
			{ pattern: /how to (.*)/i, type: 'howto' },
			{ pattern: /when (.*)/i, 'type': 'when' },
			{ pattern: /where (.*)/i, 'type': 'where' },
			{ pattern: /who is (.*)/i, type: 'person' }
		];

		for (const { pattern, type } of questionPatterns) {
			const match = lowerQuery.match(pattern);
			if (match) {
				showSpecificAnswer = true;
				specificAnswerData = {
					type,
					query: match[1] || query,
					fullQuery: query
				};
				break;
			}
		}
	}

	function handleSearch() {
		if (!searchQuery.trim()) return;
		
		isSearching = true;
		resultsLoading = true; // Start loading results for skeletons
		showSuggestions = false;
		currentPage = 1;
		
		// Detect search type and show appropriate widgets
		detectSearchType(searchQuery);
		
		setTimeout(() => {
			isSearching = false; // Main button stops spinning
			hasSearched = true; // Transition to results view
		}, 800); // Short delay for button feedback

		setTimeout(() => {
			resultsLoading = false; // Actual results appear after a longer delay
		}, 1800); // Longer delay for content loading simulation
	}

	function handleKeyPress(event) {
		if (event.key === 'Enter') {
			handleSearch();
		}
	}

	function handleInput() {
		if (searchQuery.length > 0) {
			suggestions = allSuggestions
				.filter(s => s.toLowerCase().includes(searchQuery.toLowerCase()))
				.slice(0, 6);
			showSuggestions = suggestions.length > 0;
		} else {
			showSuggestions = false;
		}
	}

	function selectSuggestion(suggestion) {
		searchQuery = suggestion;
		showSuggestions = false;
		handleSearch();
	}

	function selectTrendingSearch(query) {
		searchQuery = query;
		showSuggestions = false;
		handleSearch();
	}

	function goToPage(page) {
		currentPage = page;
		document.querySelector('main')?.scrollIntoView({ behavior: 'smooth' });
	}

	function exportResults(format) {
		showExportMenu = false;
		
		if (format === 'pdf') {
			const link = document.createElement('a');
			link.href = 'data:application/pdf;base64,JVBERi0xLjQKJdPr6eEKMSAwIG9iago8PAovVGl0bGUgKEJyZWV6ZSBTZWFyY2ggUmVzdWx0cykKL0NyZWF0b3IgKEJyZWV6ZSBTZWFyY2ggRW5naW5lKQovUHJvZHVjZXIgKEJyZWV6ZSBTZWFyY2ggRW5naW5lKQovQ3JlYXRpb25EYXRlIChEOjIwMjQwMTAxMTIwMDAwWikKPj4KZW5kb2JqCjIgMCBvYmoKPDwKL1R5cGUgL0NhdGFsb2cKL1BhZ2VzIDMgMCBSCj4+CmVuZG9iagozIDAgb2JqCjw8Ci9UeXBlIC9QYWdlcwovS2lkcyBbNCAwIFJdCi9Db3VudCAxCj4+CmVuZG9iago0IDAgb2JqCjw8Ci9UeXBlIC9QYWdlCi9QYXJlbnQgMyAwIFIKL01lZGlhQm94IFswIDAgNjEyIDc5Ml0KL0NvbnRlbnRzIDUgMCBSCj4+CmVuZG9iago1IDAgb2JqCjw8Ci9MZW5ndGggNDQKPj4Kc3RyZWFtCkJUCi9GMSAxMiBUZgoxMDAgNzAwIFRkCihCcmVlemUgU2VhcmNoIFJlc3VsdHMpIFRqCkVUCmVuZHN0cmVhbQplbmRvYmoKeHJlZgowIDYKMDAwMDAwMDAwMCA2NTUzNSBmIAowMDAwMDAwMDA5IDAwMDAwIG4gCjAwMDAwMDAxNzQgMDAwMDAgbiAKMDAwMDAwMDIyMSAwMDAwMCBuIAowMDAwMDAwMjc4IDAwMDAwIG4gCjAwMDAwMDAzNzUgMDAwMDAgbiAKdHJhaWxlcgo8PAovU2l6ZSA2Ci9Sb290IDIgMCBSCj4+CnN0YXJ0eHJlZgo0NjkKJSVFT0Y=';
			link.download = `breeze-search-results-${searchQuery.replace(/\s+/g, '-')}.pdf`;
			link.click();
		} else if (format === 'csv') {
			const csvContent = [
				'Title,URL,Description,Timestamp,Verified',
				...mockResults.map(result => 
					`"${result.title}","${result.url}","${result.description}","${result.timestamp}","${result.verified}"`
				)
			].join('\n');
			
			const blob = new Blob([csvContent], { type: 'text/csv' });
			const link = document.createElement('a');
			link.href = URL.createObjectURL(blob);
			link.download = `breeze-search-results-${searchQuery.replace(/\s+/g, '-')}.csv`;
			link.click();
		}
	}

	// Close menus when clicking outside
	function handleClickOutside(event) {
		if (!event.target.closest('.search-container')) {
			showSuggestions = false;
		}
		if (!event.target.closest('.export-menu')) {
			showExportMenu = false;
		}
	}

	// Initialize speech recognition on mount
	$effect(() => {
		initSpeechRecognition();
	});
</script>

<svelte:window onclick={handleClickOutside} />

<div class="min-h-screen bg-gradient-to-b from-purple-950 via-purple-900 to-slate-950">
	<!-- Header -->
	<header class="sticky top-0 z-50 border-b border-purple-800/30 bg-purple-950/80 backdrop-blur-xl">
		<div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
			<div class="flex h-16 items-center justify-between">
				<div class="flex items-center gap-3">
					<div class="flex h-9 w-9 items-center justify-center rounded-xl bg-gradient-to-br from-purple-500 to-purple-700 shadow-lg shadow-purple-500/25">
						<Zap class="h-5 w-5 text-white" />
					</div>
					<div>
						<h1 class="text-xl font-semibold text-white">Breeze</h1>
						<p class="text-xs text-purple-300 -mt-0.5">Privacy-First Search</p>
					</div>
				</div>
				
				<div class="flex items-center gap-4">
					<button
						onclick={() => showPrivacyFeatures = !showPrivacyFeatures}
						class="flex items-center gap-2 rounded-full bg-purple-800/50 px-4 py-2 text-sm font-medium text-purple-200 transition-all hover:bg-purple-700/50 hover:text-white"
					>
						<Shield class="h-4 w-4" />
						Privacy Features
					</button>
				</div>
			</div>
		</div>
	</header>

	{#if showPrivacyFeatures}
		<!-- Privacy Features Banner -->
		<div class="border-b border-purple-800/30 bg-purple-900/30">
			<div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8 py-4">
				<div class="grid gap-6 sm:grid-cols-3">
					<div class="flex items-center gap-3">
						<div class="flex h-8 w-8 items-center justify-center rounded-lg bg-purple-800/50">
							<EyeOff class="h-4 w-4 text-purple-300" />
						</div>
						<div>
							<p class="text-sm font-medium text-white">No Tracking</p>
							<p class="text-xs text-purple-300">Zero data collection</p>
						</div>
					</div>
					<div class="flex items-center gap-3">
						<div class="flex h-8 w-8 items-center justify-center rounded-lg bg-purple-800/50">
							<Lock class="h-4 w-4 text-purple-300" />
						</div>
						<div>
							<p class="text-sm font-medium text-white">Encrypted</p>
							<p class="text-xs text-purple-300">End-to-end security</p>
						</div>
					</div>
					<div class="flex items-center gap-3">
						<div class="flex h-8 w-8 items-center justify-center rounded-lg bg-purple-800/50">
							<Shield class="h-4 w-4 text-purple-300" />
						</div>
						<div>
							<p class="text-sm font-medium text-white">Ad-Free</p>
							<p class="text-xs text-purple-300">Pure search results</p>
						</div>
					</div>
				</div>
			</div>
		</div>
	{/if}

	<main class="mx-auto max-w-4xl px-4 sm:px-6 lg:px-8">
		{#if !hasSearched}
			<!-- Hero Section -->
			<div class="py-20 text-center">
				<div class="mx-auto mb-8 flex h-20 w-20 items-center justify-center rounded-2xl bg-gradient-to-br from-purple-500 to-purple-700 shadow-2xl shadow-purple-500/25">
					<Zap class="h-10 w-10 text-white" />
				</div>
				<h1 class="mb-4 text-5xl font-bold tracking-tight text-white sm:text-6xl">
					Search with
					<span class="bg-gradient-to-r from-purple-400 to-purple-600 bg-clip-text text-transparent">
						confidence
					</span>
				</h1>
				<p class="mx-auto mb-12 max-w-2xl text-xl text-purple-200 leading-relaxed">
					Experience the web without compromise. No tracking, no ads, no data collection. 
					Just pure, private search results.
				</p>
			</div>
		{:else}
			<div class="py-8"></div>
		{/if}

		<!-- Search Section -->
		<div class="mb-12 search-container">
			<div class="relative">
				<div class="flex gap-3">
					<div class="relative flex-1">
						<Search class="absolute left-4 top-1/2 h-5 w-5 -translate-y-1/2 text-purple-400" />
						<Input
							bind:value={searchQuery}
							oninput={handleInput}
							onkeypress={handleKeyPress}
							onfocus={handleInput}
							placeholder="Search privately..."
							class="h-14 pl-12 pr-16 text-lg border-purple-700/50 bg-purple-900/30 text-white placeholder:text-purple-300 shadow-sm focus:border-purple-500 focus:ring-purple-500/20 focus:ring-4 transition-all duration-200"
						/>
						
						<!-- Voice Search Button -->
						<button
							onclick={isListening ? stopVoiceSearch : startVoiceSearch}
							class={`absolute right-4 top-1/2 -translate-y-1/2 h-8 w-8 rounded-lg flex items-center justify-center transition-all ${
								isListening 
									? 'bg-red-500 text-white animate-pulse' 
									: 'bg-purple-700/50 text-purple-300 hover:bg-purple-600/50 hover:text-white'
							}`}
						>
							{#if isListening}
								<MicOff class="h-4 w-4" />
							{:else}
								<Mic class="h-4 w-4" />
							{/if}
						</button>
						
						<!-- Autocomplete Suggestions -->
						{#if showSuggestions}
							<div class="absolute top-full left-0 right-0 mt-2 bg-purple-900/90 border border-purple-700/50 rounded-xl shadow-lg backdrop-blur-sm z-50">
								{#each suggestions as suggestion}
									<button
										onclick={() => selectSuggestion(suggestion)}
										class="w-full px-4 py-3 text-left text-purple-200 hover:bg-purple-800/50 hover:text-white first:rounded-t-xl last:rounded-b-xl transition-colors"
									>
										<div class="flex items-center gap-3">
											<Search class="h-4 w-4 text-purple-400" />
											<span>{suggestion}</span>
										</div>
									</button>
								{/each}
							</div>
						{/if}
					</div>
					<Button
						onclick={handleSearch}
						disabled={isSearching || !searchQuery.trim()}
						class="h-14 px-8 bg-gradient-to-r from-purple-600 to-purple-700 hover:from-purple-700 hover:to-purple-800 text-white font-medium shadow-lg shadow-purple-500/25 transition-all duration-200 hover:shadow-xl hover:shadow-purple-500/30"
					>
						{#if isSearching}
							<div class="h-5 w-5 animate-spin rounded-full border-2 border-white/30 border-t-white"></div>
						{:else}
							Search
						{/if}
					</Button>
				</div>
			</div>
		</div>

		{#if !hasSearched}
			<!-- Trending & Recent -->
			<div class="grid gap-8 lg:grid-cols-2">
				<Card class="border-purple-700/30 bg-purple-900/20 backdrop-blur-sm shadow-sm hover:shadow-md transition-shadow duration-200">
					<CardContent class="p-6">
						<div class="mb-6 flex items-center gap-3">
							<div class="flex h-8 w-8 items-center justify-center rounded-lg bg-orange-500/20">
								<TrendingUp class="h-4 w-4 text-orange-400" />
							</div>
							<h3 class="text-lg font-semibold text-white">Trending Searches</h3>
						</div>
						<div class="space-y-1">
							{#each trendingSearches as search}
								<button
									onclick={() => selectTrendingSearch(search)}
									class="block w-full rounded-lg p-3 text-left text-purple-200 transition-colors duration-150 hover:bg-purple-800/30 hover:text-white"
								>
									{search}
								</button>
							{/each}
						</div>
					</CardContent>
				</Card>

				<Card class="border-purple-700/30 bg-purple-900/20 backdrop-blur-sm shadow-sm hover:shadow-md transition-shadow duration-200">
					<CardContent class="p-6">
						<div class="mb-6 flex items-center gap-3">
							<div class="flex h-8 w-8 items-center justify-center rounded-lg bg-blue-500/20">
								<Clock class="h-4 w-4 text-blue-400" />
							</div>
							<h3 class="text-lg font-semibold text-white">Recent Searches</h3>
						</div>
						<div class="space-y-1">
							{#each recentSearches as search}
								<button
									onclick={() => selectTrendingSearch(search)}
									class="block w-full rounded-lg p-3 text-left text-purple-200 transition-colors duration-150 hover:bg-purple-800/30 hover:text-white"
								>
									{search}
								</button>
							{/each}
						</div>
					</CardContent>
				</Card>
			</div>

			<!-- Privacy Promise -->
			<div class="mt-16 rounded-2xl bg-gradient-to-r from-purple-900/30 to-purple-800/30 backdrop-blur-sm p-8 text-center border border-purple-700/30">
				<div class="mx-auto mb-4 flex h-12 w-12 items-center justify-center rounded-xl bg-gradient-to-br from-purple-500 to-purple-700">
					<Shield class="h-6 w-6 text-white" />
				</div>
				<h3 class="mb-2 text-xl font-semibold text-white">Your Privacy, Guaranteed</h3>
				<p class="mx-auto max-w-2xl text-purple-200">
					We don't store your searches, track your activity, or sell your data. 
					Search freely knowing your privacy is protected.
				</p>
			</div>
		{/if}

		{#if hasSearched}
			<!-- Search Results -->
			<div class="space-y-6">
				<div class="flex items-center justify-between border-b border-purple-700/30 pb-4">
					<div>
						<p class="text-purple-200">
							About <span class="font-medium text-white">847,000</span> results for 
							<span class="font-medium text-white">"{searchQuery}"</span>
						</p>
						<p class="text-sm text-purple-400 mt-1">Searched privately in 0.31 seconds</p>
					</div>
					<div class="flex items-center gap-3">
						<Badge variant="secondary" class="bg-green-500/20 text-green-400 border-green-500/30">
							<Shield class="h-3 w-3 mr-1" />
							Private
						</Badge>
						
						<!-- Export Menu -->
						<div class="relative export-menu">
							<Button
								variant="outline"
								size="sm"
								onclick={() => showExportMenu = !showExportMenu}
								class="border-purple-600 text-purple-300 hover:bg-purple-800/30 hover:text-white"
							>
								<Download class="h-4 w-4 mr-2" />
								Export
								<ChevronDown class="h-3 w-3 ml-2" />
							</Button>
							
							{#if showExportMenu}
								<div class="absolute right-0 top-full mt-2 w-48 bg-purple-900/90 border border-purple-700/50 rounded-lg shadow-lg backdrop-blur-sm z-50">
									<button
										onclick={() => exportResults('pdf')}
										class="w-full px-4 py-3 text-left text-purple-200 hover:bg-purple-800/50 hover:text-white first:rounded-t-lg transition-colors flex items-center gap-3"
									>
										<FileText class="h-4 w-4 text-red-400" />
										<div>
											<p class="font-medium">Export as PDF</p>
											<p class="text-xs text-purple-400">Formatted document</p>
										</div>
									</button>
									<button
										onclick={() => exportResults('csv')}
										class="w-full px-4 py-3 text-left text-purple-200 hover:bg-purple-800/50 hover:text-white last:rounded-b-lg transition-colors flex items-center gap-3"
									>
										<FileSpreadsheet class="h-4 w-4 text-green-400" />
										<div>
											<p class="font-medium">Export as CSV</p>
											<p class="text-xs text-purple-400">Spreadsheet data</p>
										</div>
									</button>
								</div>
							{/if}
						</div>
					</div>
				</div>

				<!-- Conditional Widgets -->
				{#if resultsLoading}
					<SearchResultSkeleton />
					<SearchResultSkeleton />
					<SearchResultSkeleton />
				{:else}
					{#if showSpecificAnswer}
						<SpecificAnswer data={specificAnswerData} />
					{/if}

					{#if showCurrencyExchange}
						<CurrencyExchange query={searchQuery} />
					{/if}

					{#if showWeather}
						<WeatherWidget query={searchQuery} />
					{/if}

					{#if showSolanaExplorer}
						<SolanaExplorer query={solanaQuery} queryType={solanaQueryType} />
					{/if}

					{#if showTokenPriceChart}
						<TokenPriceChart tokenSymbol={tokenChartSymbol} />
					{/if}

					<!-- AI Summary -->
					<AISummary {searchQuery} />

					<div class="space-y-6">
						{#each mockResults as result, index}
							<Card class="border-purple-700/30 bg-purple-900/20 backdrop-blur-sm shadow-sm hover:shadow-md transition-all duration-200 hover:border-purple-600/50">
								<CardContent class="p-6">
									<div class="space-y-3">
										<div class="flex items-start gap-4">
											<SiteLogo domain={result.domain} hasLogo={result.hasLogo} />
											<div class="flex-1 min-w-0">
												<div class="flex items-center gap-2 mb-1">
													<a href={'https://' + result.url} class="text-sm text-purple-400 hover:text-purple-300 truncate">
														{result.url}
													</a>
													{#if result.verified}
														<Badge variant="outline" class="border-emerald-500/30 text-emerald-400 text-xs flex-shrink-0">
															<Shield class="h-2.5 w-2.5 mr-1" />
															Verified
														</Badge>
													{/if}
													<span class="text-xs text-purple-400 ml-auto flex-shrink-0">
														{result.timestamp}
													</span>
												</div>
												<h3 class="text-xl font-semibold text-white hover:text-purple-400 cursor-pointer transition-colors mb-2">
													{result.title}
												</h3>
												<p class="text-purple-200 leading-relaxed">
													{result.description}
												</p>
											</div>
										</div>
									</div>
								</CardContent>
							</Card>

							<!-- Insert News Carousel after first 2 results -->
							{#if index === 1}
								<NewsCarousel />
							{/if}
						{/each}
					</div>

					<!-- Pagination -->
					<div class="py-8">
						<div class="flex items-center justify-center gap-2">
							<Button
								variant="outline"
								size="sm"
								disabled={currentPage === 1}
								onclick={() => goToPage(currentPage - 1)}
								class="border-purple-600 text-purple-300 hover:bg-purple-800/30 hover:text-white"
							>
								<ChevronLeft class="h-4 w-4" />
							</Button>
							
							{#each Array.from({length: Math.min(7, totalPages)}, (_, i) => {
								const start = Math.max(1, currentPage - 3);
								return start + i;
							}) as page}
								{#if page <= totalPages}
									<Button
										variant={currentPage === page ? "default" : "outline"}
										size="sm"
										onclick={() => goToPage(page)}
										class={currentPage === page 
											? "bg-gradient-to-r from-purple-600 to-purple-700 text-white" 
											: "border-purple-600 text-purple-300 hover:bg-purple-800/30 hover:text-white"
										}
									>
										{page}
									</Button>
								{/if}
							{/each}
							
							<Button
								variant="outline"
								size="sm"
								disabled={currentPage === totalPages}
								onclick={() => goToPage(currentPage + 1)}
								class="border-purple-600 text-purple-300 hover:bg-purple-800/30 hover:text-white"
							>
								<ChevronRight class="h-4 w-4" />
							</Button>
						</div>
						<p class="text-center text-sm text-purple-400 mt-4">
							Page {currentPage} of {totalPages} • Showing {(currentPage - 1) * resultsPerPage + 1}-{Math.min(currentPage * resultsPerPage, 847000)} results
						</p>
					</div>
				{/if}
			</div>
		{/if}
	</main>

	<!-- Footer -->
	<footer class="mt-20 border-t border-purple-800/30 bg-purple-950/30">
		<div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8 py-12">
			<div class="text-center">
				<div class="mb-4 flex items-center justify-center gap-2">
					<div class="flex h-6 w-6 items-center justify-center rounded-lg bg-gradient-to-br from-purple-500 to-purple-700">
						<Zap class="h-3 w-3 text-white" />
					</div>
					<span class="font-semibold text-white">Breeze Search</span>
				</div>
				<p class="text-purple-200 mb-4">
					Privacy-first search engine. No tracking, no ads, no compromise.
				</p>
				<p class="text-sm text-purple-400">
					© 2024 Breeze Search. Built with privacy in mind.
				</p>
			</div>
		</div>
	</footer>
</div>
