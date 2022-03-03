package public

import "embed"

// Public is our viewsFS web server content
//go:embed build/*
var Public embed.FS
