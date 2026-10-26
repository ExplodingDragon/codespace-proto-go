// Copyright 2026 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

// Package codespacev1 contains the shared protocol surface for the codespace system.
package codespacev1

// ManagerServiceMaxMessageBytes bounds both requests and responses independently of runtime storage limits.
const ManagerServiceMaxMessageBytes = 32 * 1024 * 1024
