// Generated from FQL.g4 by ANTLR 4.7.

package fql

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 27, 356,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3,
	12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16,
	3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 184, 10, 17, 3, 18, 3,
	18, 3, 18, 3, 18, 5, 18, 190, 10, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21,
	3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3,
	26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31,
	3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3, 36, 3, 36, 3,
	37, 3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 41, 3, 41, 3, 42,
	3, 42, 3, 43, 3, 43, 3, 44, 3, 44, 3, 45, 3, 45, 3, 46, 3, 46, 3, 47, 3,
	47, 3, 47, 3, 47, 3, 48, 3, 48, 7, 48, 254, 10, 48, 12, 48, 14, 48, 257,
	11, 48, 3, 49, 6, 49, 260, 10, 49, 13, 49, 14, 49, 261, 3, 50, 6, 50, 265,
	10, 50, 13, 50, 14, 50, 266, 3, 50, 3, 50, 7, 50, 271, 10, 50, 12, 50,
	14, 50, 274, 11, 50, 3, 50, 5, 50, 277, 10, 50, 3, 50, 3, 50, 6, 50, 281,
	10, 50, 13, 50, 14, 50, 282, 3, 50, 5, 50, 286, 10, 50, 3, 50, 6, 50, 289,
	10, 50, 13, 50, 14, 50, 290, 3, 50, 5, 50, 294, 10, 50, 3, 51, 3, 51, 3,
	51, 7, 51, 299, 10, 51, 12, 51, 14, 51, 302, 11, 51, 3, 51, 3, 51, 3, 52,
	3, 52, 3, 52, 5, 52, 309, 10, 52, 3, 53, 3, 53, 3, 53, 5, 53, 314, 10,
	53, 3, 53, 3, 53, 3, 54, 3, 54, 5, 54, 320, 10, 54, 3, 54, 6, 54, 323,
	10, 54, 13, 54, 14, 54, 324, 3, 55, 3, 55, 3, 56, 3, 56, 3, 56, 3, 56,
	5, 56, 333, 10, 56, 3, 56, 3, 56, 5, 56, 337, 10, 56, 3, 57, 3, 57, 3,
	57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 5, 57, 348, 10, 57, 3, 58,
	3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 2, 2, 59, 3, 3, 5, 4, 7, 5, 9,
	6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15,
	29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 2, 43, 2, 45, 2, 47,
	2, 49, 2, 51, 2, 53, 2, 55, 2, 57, 2, 59, 2, 61, 2, 63, 2, 65, 2, 67, 2,
	69, 2, 71, 2, 73, 2, 75, 2, 77, 2, 79, 2, 81, 2, 83, 2, 85, 2, 87, 2, 89,
	2, 91, 2, 93, 22, 95, 23, 97, 24, 99, 25, 101, 26, 103, 2, 105, 27, 107,
	2, 109, 2, 111, 2, 113, 2, 115, 2, 3, 2, 36, 4, 2, 67, 67, 99, 99, 4, 2,
	68, 68, 100, 100, 4, 2, 69, 69, 101, 101, 4, 2, 70, 70, 102, 102, 4, 2,
	71, 71, 103, 103, 4, 2, 72, 72, 104, 104, 4, 2, 73, 73, 105, 105, 4, 2,
	74, 74, 106, 106, 4, 2, 75, 75, 107, 107, 4, 2, 76, 76, 108, 108, 4, 2,
	77, 77, 109, 109, 4, 2, 78, 78, 110, 110, 4, 2, 79, 79, 111, 111, 4, 2,
	80, 80, 112, 112, 4, 2, 81, 81, 113, 113, 4, 2, 82, 82, 114, 114, 4, 2,
	83, 83, 115, 115, 4, 2, 84, 84, 116, 116, 4, 2, 85, 85, 117, 117, 4, 2,
	86, 86, 118, 118, 4, 2, 87, 87, 119, 119, 4, 2, 88, 88, 120, 120, 4, 2,
	89, 89, 121, 121, 4, 2, 90, 90, 122, 122, 4, 2, 91, 91, 123, 123, 4, 2,
	92, 92, 124, 124, 5, 2, 11, 12, 15, 15, 34, 34, 6, 2, 49, 49, 67, 92, 97,
	97, 99, 124, 6, 2, 47, 59, 67, 92, 97, 97, 99, 124, 4, 2, 36, 36, 94, 94,
	4, 2, 41, 41, 94, 94, 4, 2, 45, 45, 47, 47, 5, 2, 50, 59, 67, 72, 99, 104,
	7, 2, 100, 100, 104, 104, 112, 112, 116, 116, 118, 118, 2, 345, 2, 3, 3,
	2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3,
	2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19,
	3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2,
	27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2,
	2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 93, 3, 2, 2,
	2, 2, 95, 3, 2, 2, 2, 2, 97, 3, 2, 2, 2, 2, 99, 3, 2, 2, 2, 2, 101, 3,
	2, 2, 2, 2, 105, 3, 2, 2, 2, 3, 117, 3, 2, 2, 2, 5, 121, 3, 2, 2, 2, 7,
	128, 3, 2, 2, 2, 9, 135, 3, 2, 2, 2, 11, 142, 3, 2, 2, 2, 13, 146, 3, 2,
	2, 2, 15, 153, 3, 2, 2, 2, 17, 158, 3, 2, 2, 2, 19, 160, 3, 2, 2, 2, 21,
	162, 3, 2, 2, 2, 23, 165, 3, 2, 2, 2, 25, 167, 3, 2, 2, 2, 27, 170, 3,
	2, 2, 2, 29, 172, 3, 2, 2, 2, 31, 175, 3, 2, 2, 2, 33, 183, 3, 2, 2, 2,
	35, 189, 3, 2, 2, 2, 37, 191, 3, 2, 2, 2, 39, 193, 3, 2, 2, 2, 41, 195,
	3, 2, 2, 2, 43, 197, 3, 2, 2, 2, 45, 199, 3, 2, 2, 2, 47, 201, 3, 2, 2,
	2, 49, 203, 3, 2, 2, 2, 51, 205, 3, 2, 2, 2, 53, 207, 3, 2, 2, 2, 55, 209,
	3, 2, 2, 2, 57, 211, 3, 2, 2, 2, 59, 213, 3, 2, 2, 2, 61, 215, 3, 2, 2,
	2, 63, 217, 3, 2, 2, 2, 65, 219, 3, 2, 2, 2, 67, 221, 3, 2, 2, 2, 69, 223,
	3, 2, 2, 2, 71, 225, 3, 2, 2, 2, 73, 227, 3, 2, 2, 2, 75, 229, 3, 2, 2,
	2, 77, 231, 3, 2, 2, 2, 79, 233, 3, 2, 2, 2, 81, 235, 3, 2, 2, 2, 83, 237,
	3, 2, 2, 2, 85, 239, 3, 2, 2, 2, 87, 241, 3, 2, 2, 2, 89, 243, 3, 2, 2,
	2, 91, 245, 3, 2, 2, 2, 93, 247, 3, 2, 2, 2, 95, 251, 3, 2, 2, 2, 97, 259,
	3, 2, 2, 2, 99, 293, 3, 2, 2, 2, 101, 295, 3, 2, 2, 2, 103, 305, 3, 2,
	2, 2, 105, 310, 3, 2, 2, 2, 107, 317, 3, 2, 2, 2, 109, 326, 3, 2, 2, 2,
	111, 336, 3, 2, 2, 2, 113, 347, 3, 2, 2, 2, 115, 349, 3, 2, 2, 2, 117,
	118, 5, 73, 37, 2, 118, 119, 5, 69, 35, 2, 119, 120, 5, 77, 39, 2, 120,
	4, 3, 2, 2, 2, 121, 122, 5, 77, 39, 2, 122, 123, 5, 49, 25, 2, 123, 124,
	5, 63, 32, 2, 124, 125, 5, 49, 25, 2, 125, 126, 5, 45, 23, 2, 126, 127,
	5, 79, 40, 2, 127, 6, 3, 2, 2, 2, 128, 129, 5, 57, 29, 2, 129, 130, 5,
	67, 34, 2, 130, 131, 5, 77, 39, 2, 131, 132, 5, 49, 25, 2, 132, 133, 5,
	75, 38, 2, 133, 134, 5, 79, 40, 2, 134, 8, 3, 2, 2, 2, 135, 136, 5, 47,
	24, 2, 136, 137, 5, 49, 25, 2, 137, 138, 5, 63, 32, 2, 138, 139, 5, 49,
	25, 2, 139, 140, 5, 79, 40, 2, 140, 141, 5, 49, 25, 2, 141, 10, 3, 2, 2,
	2, 142, 143, 5, 77, 39, 2, 143, 144, 5, 49, 25, 2, 144, 145, 5, 79, 40,
	2, 145, 12, 3, 2, 2, 2, 146, 147, 5, 49, 25, 2, 147, 148, 5, 87, 44, 2,
	148, 149, 5, 71, 36, 2, 149, 150, 5, 69, 35, 2, 150, 151, 5, 75, 38, 2,
	151, 152, 5, 79, 40, 2, 152, 14, 3, 2, 2, 2, 153, 154, 5, 51, 26, 2, 154,
	155, 5, 75, 38, 2, 155, 156, 5, 69, 35, 2, 156, 157, 5, 65, 33, 2, 157,
	16, 3, 2, 2, 2, 158, 159, 7, 44, 2, 2, 159, 18, 3, 2, 2, 2, 160, 161, 7,
	63, 2, 2, 161, 20, 3, 2, 2, 2, 162, 163, 7, 63, 2, 2, 163, 164, 7, 63,
	2, 2, 164, 22, 3, 2, 2, 2, 165, 166, 7, 62, 2, 2, 166, 24, 3, 2, 2, 2,
	167, 168, 7, 62, 2, 2, 168, 169, 7, 63, 2, 2, 169, 26, 3, 2, 2, 2, 170,
	171, 7, 64, 2, 2, 171, 28, 3, 2, 2, 2, 172, 173, 7, 64, 2, 2, 173, 174,
	7, 63, 2, 2, 174, 30, 3, 2, 2, 2, 175, 176, 7, 35, 2, 2, 176, 177, 7, 63,
	2, 2, 177, 32, 3, 2, 2, 2, 178, 179, 5, 41, 21, 2, 179, 180, 5, 67, 34,
	2, 180, 181, 5, 47, 24, 2, 181, 184, 3, 2, 2, 2, 182, 184, 7, 40, 2, 2,
	183, 178, 3, 2, 2, 2, 183, 182, 3, 2, 2, 2, 184, 34, 3, 2, 2, 2, 185, 186,
	5, 69, 35, 2, 186, 187, 5, 75, 38, 2, 187, 190, 3, 2, 2, 2, 188, 190, 7,
	126, 2, 2, 189, 185, 3, 2, 2, 2, 189, 188, 3, 2, 2, 2, 190, 36, 3, 2, 2,
	2, 191, 192, 7, 46, 2, 2, 192, 38, 3, 2, 2, 2, 193, 194, 7, 61, 2, 2, 194,
	40, 3, 2, 2, 2, 195, 196, 9, 2, 2, 2, 196, 42, 3, 2, 2, 2, 197, 198, 9,
	3, 2, 2, 198, 44, 3, 2, 2, 2, 199, 200, 9, 4, 2, 2, 200, 46, 3, 2, 2, 2,
	201, 202, 9, 5, 2, 2, 202, 48, 3, 2, 2, 2, 203, 204, 9, 6, 2, 2, 204, 50,
	3, 2, 2, 2, 205, 206, 9, 7, 2, 2, 206, 52, 3, 2, 2, 2, 207, 208, 9, 8,
	2, 2, 208, 54, 3, 2, 2, 2, 209, 210, 9, 9, 2, 2, 210, 56, 3, 2, 2, 2, 211,
	212, 9, 10, 2, 2, 212, 58, 3, 2, 2, 2, 213, 214, 9, 11, 2, 2, 214, 60,
	3, 2, 2, 2, 215, 216, 9, 12, 2, 2, 216, 62, 3, 2, 2, 2, 217, 218, 9, 13,
	2, 2, 218, 64, 3, 2, 2, 2, 219, 220, 9, 14, 2, 2, 220, 66, 3, 2, 2, 2,
	221, 222, 9, 15, 2, 2, 222, 68, 3, 2, 2, 2, 223, 224, 9, 16, 2, 2, 224,
	70, 3, 2, 2, 2, 225, 226, 9, 17, 2, 2, 226, 72, 3, 2, 2, 2, 227, 228, 9,
	18, 2, 2, 228, 74, 3, 2, 2, 2, 229, 230, 9, 19, 2, 2, 230, 76, 3, 2, 2,
	2, 231, 232, 9, 20, 2, 2, 232, 78, 3, 2, 2, 2, 233, 234, 9, 21, 2, 2, 234,
	80, 3, 2, 2, 2, 235, 236, 9, 22, 2, 2, 236, 82, 3, 2, 2, 2, 237, 238, 9,
	23, 2, 2, 238, 84, 3, 2, 2, 2, 239, 240, 9, 24, 2, 2, 240, 86, 3, 2, 2,
	2, 241, 242, 9, 25, 2, 2, 242, 88, 3, 2, 2, 2, 243, 244, 9, 26, 2, 2, 244,
	90, 3, 2, 2, 2, 245, 246, 9, 27, 2, 2, 246, 92, 3, 2, 2, 2, 247, 248, 9,
	28, 2, 2, 248, 249, 3, 2, 2, 2, 249, 250, 8, 47, 2, 2, 250, 94, 3, 2, 2,
	2, 251, 255, 9, 29, 2, 2, 252, 254, 9, 30, 2, 2, 253, 252, 3, 2, 2, 2,
	254, 257, 3, 2, 2, 2, 255, 253, 3, 2, 2, 2, 255, 256, 3, 2, 2, 2, 256,
	96, 3, 2, 2, 2, 257, 255, 3, 2, 2, 2, 258, 260, 4, 50, 59, 2, 259, 258,
	3, 2, 2, 2, 260, 261, 3, 2, 2, 2, 261, 259, 3, 2, 2, 2, 261, 262, 3, 2,
	2, 2, 262, 98, 3, 2, 2, 2, 263, 265, 4, 50, 59, 2, 264, 263, 3, 2, 2, 2,
	265, 266, 3, 2, 2, 2, 266, 264, 3, 2, 2, 2, 266, 267, 3, 2, 2, 2, 267,
	268, 3, 2, 2, 2, 268, 272, 7, 48, 2, 2, 269, 271, 4, 50, 59, 2, 270, 269,
	3, 2, 2, 2, 271, 274, 3, 2, 2, 2, 272, 270, 3, 2, 2, 2, 272, 273, 3, 2,
	2, 2, 273, 276, 3, 2, 2, 2, 274, 272, 3, 2, 2, 2, 275, 277, 5, 107, 54,
	2, 276, 275, 3, 2, 2, 2, 276, 277, 3, 2, 2, 2, 277, 294, 3, 2, 2, 2, 278,
	280, 7, 48, 2, 2, 279, 281, 4, 50, 59, 2, 280, 279, 3, 2, 2, 2, 281, 282,
	3, 2, 2, 2, 282, 280, 3, 2, 2, 2, 282, 283, 3, 2, 2, 2, 283, 285, 3, 2,
	2, 2, 284, 286, 5, 107, 54, 2, 285, 284, 3, 2, 2, 2, 285, 286, 3, 2, 2,
	2, 286, 294, 3, 2, 2, 2, 287, 289, 4, 50, 59, 2, 288, 287, 3, 2, 2, 2,
	289, 290, 3, 2, 2, 2, 290, 288, 3, 2, 2, 2, 290, 291, 3, 2, 2, 2, 291,
	292, 3, 2, 2, 2, 292, 294, 5, 107, 54, 2, 293, 264, 3, 2, 2, 2, 293, 278,
	3, 2, 2, 2, 293, 288, 3, 2, 2, 2, 294, 100, 3, 2, 2, 2, 295, 300, 7, 36,
	2, 2, 296, 299, 5, 103, 52, 2, 297, 299, 10, 31, 2, 2, 298, 296, 3, 2,
	2, 2, 298, 297, 3, 2, 2, 2, 299, 302, 3, 2, 2, 2, 300, 298, 3, 2, 2, 2,
	300, 301, 3, 2, 2, 2, 301, 303, 3, 2, 2, 2, 302, 300, 3, 2, 2, 2, 303,
	304, 7, 36, 2, 2, 304, 102, 3, 2, 2, 2, 305, 308, 7, 94, 2, 2, 307, 309,
	9, 32, 2, 2, 308, 306, 3, 2, 2, 2, 308, 307, 3, 2, 2, 2, 309, 104, 3, 2,
	2, 2, 310, 313, 7, 41, 2, 2, 311, 314, 5, 111, 56, 2, 312, 314, 10, 32,
	2, 2, 313, 311, 3, 2, 2, 2, 313, 312, 3, 2, 2, 2, 314, 315, 3, 2, 2, 2,
	315, 316, 7, 41, 2, 2, 316, 106, 3, 2, 2, 2, 317, 319, 9, 6, 2, 2, 318,
	320, 9, 33, 2, 2, 319, 318, 3, 2, 2, 2, 319, 320, 3, 2, 2, 2, 320, 322,
	3, 2, 2, 2, 321, 323, 4, 50, 59, 2, 322, 321, 3, 2, 2, 2, 323, 324, 3,
	2, 2, 2, 324, 322, 3, 2, 2, 2, 324, 325, 3, 2, 2, 2, 325, 108, 3, 2, 2,
	2, 326, 327, 9, 34, 2, 2, 327, 110, 3, 2, 2, 2, 328, 332, 7, 94, 2, 2,
	329, 333, 9, 35, 2, 2, 331, 333, 9, 32, 2, 2, 332, 329, 3, 2, 2, 2, 332,
	330, 3, 2, 2, 2, 332, 331, 3, 2, 2, 2, 333, 337, 3, 2, 2, 2, 334, 337,
	5, 115, 58, 2, 335, 337, 5, 113, 57, 2, 336, 328, 3, 2, 2, 2, 336, 334,
	3, 2, 2, 2, 336, 335, 3, 2, 2, 2, 337, 112, 3, 2, 2, 2, 338, 339, 7, 94,
	2, 2, 339, 340, 4, 50, 53, 2, 340, 341, 4, 50, 57, 2, 341, 348, 4, 50,
	57, 2, 342, 343, 7, 94, 2, 2, 343, 344, 4, 50, 57, 2, 344, 348, 4, 50,
	57, 2, 345, 346, 7, 94, 2, 2, 346, 348, 4, 50, 57, 2, 347, 338, 3, 2, 2,
	2, 347, 342, 3, 2, 2, 2, 347, 345, 3, 2, 2, 2, 348, 114, 3, 2, 2, 2, 349,
	350, 7, 94, 2, 2, 350, 351, 7, 119, 2, 2, 351, 352, 5, 109, 55, 2, 352,
	353, 5, 109, 55, 2, 353, 354, 5, 109, 55, 2, 354, 355, 5, 109, 55, 2, 355,
	116, 3, 2, 2, 2, 23, 2, 183, 189, 255, 261, 266, 272, 276, 282, 285, 290,
	293, 298, 300, 308, 313, 319, 324, 332, 336, 347, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "", "", "", "", "", "", "", "'*'", "'='", "'=='", "'<'", "'<='", "'>'",
	"'>='", "'!='", "", "", "','", "';'",
}

var lexerSymbolicNames = []string{
	"", "QOS", "SELECT", "INSERT", "DELETE", "SET", "EXPORT", "FROM", "ASTERISK",
	"SINGLE_EQ", "DOUBLE_EQ", "OP_LT", "LE", "GT", "GE", "NOTEQ", "AND", "OR",
	"COMMA", "SEMICOLON", "WS", "ID", "NUMBER", "FLOAT", "STRING", "CHAR",
}

var lexerRuleNames = []string{
	"QOS", "SELECT", "INSERT", "DELETE", "SET", "EXPORT", "FROM", "ASTERISK",
	"SINGLE_EQ", "DOUBLE_EQ", "OP_LT", "LE", "GT", "GE", "NOTEQ", "AND", "OR",
	"COMMA", "SEMICOLON", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
	"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y",
	"Z", "WS", "ID", "NUMBER", "FLOAT", "STRING", "EscapeSequence", "CHAR",
	"EXPONENT", "HEX_DIGIT", "ESC_SEQ", "OCTAL_ESC", "UNICODE_ESC",
}

type FQLLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewFQLLexer(input antlr.CharStream) *FQLLexer {

	l := new(FQLLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "FQL.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// FQLLexer tokens.
const (
	FQLLexerQOS       = 1
	FQLLexerSELECT    = 2
	FQLLexerINSERT    = 3
	FQLLexerDELETE    = 4
	FQLLexerSET       = 5
	FQLLexerEXPORT    = 6
	FQLLexerFROM      = 7
	FQLLexerASTERISK  = 8
	FQLLexerSINGLE_EQ = 9
	FQLLexerDOUBLE_EQ = 10
	FQLLexerOP_LT     = 11
	FQLLexerLE        = 12
	FQLLexerGT        = 13
	FQLLexerGE        = 14
	FQLLexerNOTEQ     = 15
	FQLLexerAND       = 16
	FQLLexerOR        = 17
	FQLLexerCOMMA     = 18
	FQLLexerSEMICOLON = 19
	FQLLexerWS        = 20
	FQLLexerID        = 21
	FQLLexerNUMBER    = 22
	FQLLexerFLOAT     = 23
	FQLLexerSTRING    = 24
	FQLLexerCHAR      = 25
)
