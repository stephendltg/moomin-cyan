#!/usr/bin/env deno run --allow-env --allow-run --allow-read --allow-write

import * as log from "https://deno.land/std@0.104.0/log/mod.ts";
import { copy, ensureDir } from "https://deno.land/std@0.104.0/fs/mod.ts";

const data = JSON.parse(Deno.readTextFileSync("./package.json"));

const template = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>CFBundleExecutable</key>
	<string>${data.name}</string>
	<key>CFBundleIconFile</key>
	<string>icon</string>
	<key>CFBundleIdentifier</key>
	<string>com.${data.name}.1</string>
	<key>NSHighResolutionCapable</key>
	<true/>
	<key>LSUIElement</key>
	<true/>
	<key>NSAppTransportSecurity</key>
	<dict>
    	<key>NSAllowsLocalNetworking</key>
    	<true/>
	</dict>
	<!-- <dict>
		<key>NSAllowsArbitraryLoads</key>
		<true/>
	</dict> -->
</dict>
</plist>
`;

var args;

if (Deno.build.os === "windows") {
  args = [
    "go",
    "build",
    '-ldflags="-H windowsgui"',
    "-v",
    "-o",
    `bin/${data.name}-win32-amd64.exe`,
  ];
} else if (Deno.build.os === "darwin") {
  // Prepare package
  await ensureDir(`./bin/${data.name}.app/Contents/MacOS`);
  await ensureDir(`./bin/${data.name}.app/Contents/Resources`);
  await copy(
    "./assets/icon.icns",
    `./bin/${data.name}.app/Contents/Resources/icon.icns`,
    { overwrite: true },
  );
  await Deno.writeTextFile(
    `./bin/${data.name}.app/Contents/Info.plist`,
    template,
  );
  args = [
    "go",
    "build",
    "-v",
    "-o",
    `bin/${data.name}.app/Contents/MacOS/${data.name}`,
  ];
} else if (Deno.build.os === "linux") {
  args = ["make", "build-deb"];
} else {
  log.critical("⚠ Unsupported platform: " + Deno.build.os);
}

if (!args) {
  Deno.exit(0);
}

// create go instance
const go = Deno.run({
  cmd: args,
  stdout: "inherit",
  stderr: "piped",
});

// await its completion
const { code } = await go.status();
const rawError = await go.stderrOutput();

if (code === 0) {
  log.info("go build ...");
} else {
  const errorString = new TextDecoder().decode(rawError);
  log.warning(errorString);
}

Deno.exit(code);
