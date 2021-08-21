#!/usr/bin/env deno run --allow-env --allow-run --allow-read

import * as log from "https://deno.land/std@0.104.0/log/mod.ts";

var goBuildArgs;

if (Deno.build.os === "linux") {
  goBuildArgs = ["GOOS=linux", "GOARCH=amd64", "go", "build", "-v", "-o", "bin/moomin-linux-amd64", "." ];  
} else {
  log.critical("⚠ Unsupported platform: " + Deno.build.os);
}

if (!goBuildArgs) {
  Deno.exit(0);
}

// create go instance
const go = Deno.run({
  cmd: goBuildArgs,
  stdout: "piped",
  stderr: "piped",
});

// await its completion
const { code } = await go.status();
const rawOutput = await go.output();
const rawError = await go.stderrOutput();

if (code === 0) {
  await Deno.stdout.write(rawOutput);
} else {
  const errorString = new TextDecoder().decode(rawError);
  log.warning(errorString);
}

Deno.exit(code);
