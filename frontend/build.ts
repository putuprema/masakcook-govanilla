import { rm } from "node:fs/promises";
import { Glob } from "bun";
import tailwind from "bun-plugin-tailwind";

const isDevelopment = Bun.env.ENVIRONMENT !== "production";

console.log(
	`Building client-side code (${isDevelopment ? "development" : "production"})...`,
);

try {
	console.log("Cleaning static build output directory...");
	await rm("../static/", { recursive: true, force: true });

	const pageFiles = await findPageFiles();
	const pageCssFiles = await findPageCssFiles();

	// Combine common.ts with all page files
	const entrypoints = [
		"./scripts/global.ts",
		"./styles/global.css",
		...pageFiles,
		...pageCssFiles,
	];

	console.log("Entry points:");
	for (const entry of entrypoints) {
		console.log(`  • ${entry}`);
	}

	const buildResult = await Bun.build({
		entrypoints,
		plugins: [tailwind],
		outdir: "../static",
		target: "browser",
		format: "esm",
		minify: !isDevelopment,
		sourcemap: isDevelopment ? "linked" : "none",
		naming: {
			entry: "[dir]/[name].[ext]",
			chunk: "[dir]/[name]-[hash].[ext]",
		},
		splitting: true,
	});

	if (!buildResult.success) {
		console.error("Build failed:");
		for (const log of buildResult.logs) {
			console.error(log);
		}
		process.exit(1);
	}

	console.log("✓ Build complete!");
	console.log("Outputs:");
	for (const output of buildResult.outputs) {
		const size = (output.size / 1024).toFixed(2);
		console.log(`  → ${output.path} (${size} KB)`);
	}
} catch (error) {
	console.error("Build error:", error);
	process.exit(1);
}

/**
 * Find all page TypeScript files
 */
async function findPageFiles(): Promise<string[]> {
	const glob = new Glob("scripts/pages/*.ts");
	const files: string[] = [];

	for await (const file of glob.scan(".")) {
		files.push(`./${file}`);
	}

	return files;
}

/**
 * Find all page CSS files
 */
async function findPageCssFiles(): Promise<string[]> {
	const glob = new Glob("styles/pages/*.css");
	const files: string[] = [];

	for await (const file of glob.scan(".")) {
		files.push(`./${file}`);
	}

	return files;
}
