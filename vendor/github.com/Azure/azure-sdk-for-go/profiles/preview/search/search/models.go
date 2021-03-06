// +build go1.9

// Copyright 2017 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder
// commit ID: 2014fbbf031942474ad27a5a66dffaed5347f3fb

package search

import original "github.com/Azure/azure-sdk-for-go/services/search/2016-09-01/search"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type DataSourcesClient = original.DataSourcesClient
type DocumentsProxyClient = original.DocumentsProxyClient
type IndexersClient = original.IndexersClient
type IndexesClient = original.IndexesClient
type CjkBigramTokenFilterScripts = original.CjkBigramTokenFilterScripts

const (
	Han      CjkBigramTokenFilterScripts = original.Han
	Hangul   CjkBigramTokenFilterScripts = original.Hangul
	Hiragana CjkBigramTokenFilterScripts = original.Hiragana
	Katakana CjkBigramTokenFilterScripts = original.Katakana
)

type EdgeNGramTokenFilterSide = original.EdgeNGramTokenFilterSide

const (
	Back  EdgeNGramTokenFilterSide = original.Back
	Front EdgeNGramTokenFilterSide = original.Front
)

type IndexActionType = original.IndexActionType

const (
	Delete        IndexActionType = original.Delete
	Merge         IndexActionType = original.Merge
	MergeOrUpload IndexActionType = original.MergeOrUpload
	Upload        IndexActionType = original.Upload
)

type IndexerExecutionStatus = original.IndexerExecutionStatus

const (
	InProgress       IndexerExecutionStatus = original.InProgress
	Reset            IndexerExecutionStatus = original.Reset
	Success          IndexerExecutionStatus = original.Success
	TransientFailure IndexerExecutionStatus = original.TransientFailure
)

type IndexerStatus = original.IndexerStatus

const (
	Error   IndexerStatus = original.Error
	Running IndexerStatus = original.Running
	Unknown IndexerStatus = original.Unknown
)

type MicrosoftStemmingTokenizerLanguage = original.MicrosoftStemmingTokenizerLanguage

const (
	Arabic              MicrosoftStemmingTokenizerLanguage = original.Arabic
	Bangla              MicrosoftStemmingTokenizerLanguage = original.Bangla
	Bulgarian           MicrosoftStemmingTokenizerLanguage = original.Bulgarian
	Catalan             MicrosoftStemmingTokenizerLanguage = original.Catalan
	Croatian            MicrosoftStemmingTokenizerLanguage = original.Croatian
	Czech               MicrosoftStemmingTokenizerLanguage = original.Czech
	Danish              MicrosoftStemmingTokenizerLanguage = original.Danish
	Dutch               MicrosoftStemmingTokenizerLanguage = original.Dutch
	English             MicrosoftStemmingTokenizerLanguage = original.English
	Estonian            MicrosoftStemmingTokenizerLanguage = original.Estonian
	Finnish             MicrosoftStemmingTokenizerLanguage = original.Finnish
	French              MicrosoftStemmingTokenizerLanguage = original.French
	German              MicrosoftStemmingTokenizerLanguage = original.German
	Greek               MicrosoftStemmingTokenizerLanguage = original.Greek
	Gujarati            MicrosoftStemmingTokenizerLanguage = original.Gujarati
	Hebrew              MicrosoftStemmingTokenizerLanguage = original.Hebrew
	Hindi               MicrosoftStemmingTokenizerLanguage = original.Hindi
	Hungarian           MicrosoftStemmingTokenizerLanguage = original.Hungarian
	Icelandic           MicrosoftStemmingTokenizerLanguage = original.Icelandic
	Indonesian          MicrosoftStemmingTokenizerLanguage = original.Indonesian
	Italian             MicrosoftStemmingTokenizerLanguage = original.Italian
	Kannada             MicrosoftStemmingTokenizerLanguage = original.Kannada
	Latvian             MicrosoftStemmingTokenizerLanguage = original.Latvian
	Lithuanian          MicrosoftStemmingTokenizerLanguage = original.Lithuanian
	Malay               MicrosoftStemmingTokenizerLanguage = original.Malay
	Malayalam           MicrosoftStemmingTokenizerLanguage = original.Malayalam
	Marathi             MicrosoftStemmingTokenizerLanguage = original.Marathi
	NorwegianBokmaal    MicrosoftStemmingTokenizerLanguage = original.NorwegianBokmaal
	Polish              MicrosoftStemmingTokenizerLanguage = original.Polish
	Portuguese          MicrosoftStemmingTokenizerLanguage = original.Portuguese
	PortugueseBrazilian MicrosoftStemmingTokenizerLanguage = original.PortugueseBrazilian
	Punjabi             MicrosoftStemmingTokenizerLanguage = original.Punjabi
	Romanian            MicrosoftStemmingTokenizerLanguage = original.Romanian
	Russian             MicrosoftStemmingTokenizerLanguage = original.Russian
	SerbianCyrillic     MicrosoftStemmingTokenizerLanguage = original.SerbianCyrillic
	SerbianLatin        MicrosoftStemmingTokenizerLanguage = original.SerbianLatin
	Slovak              MicrosoftStemmingTokenizerLanguage = original.Slovak
	Slovenian           MicrosoftStemmingTokenizerLanguage = original.Slovenian
	Spanish             MicrosoftStemmingTokenizerLanguage = original.Spanish
	Swedish             MicrosoftStemmingTokenizerLanguage = original.Swedish
	Tamil               MicrosoftStemmingTokenizerLanguage = original.Tamil
	Telugu              MicrosoftStemmingTokenizerLanguage = original.Telugu
	Turkish             MicrosoftStemmingTokenizerLanguage = original.Turkish
	Ukrainian           MicrosoftStemmingTokenizerLanguage = original.Ukrainian
	Urdu                MicrosoftStemmingTokenizerLanguage = original.Urdu
)

type MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguage

const (
	MicrosoftTokenizerLanguageBangla              MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguageBangla
	MicrosoftTokenizerLanguageBulgarian           MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguageBulgarian
	MicrosoftTokenizerLanguageCatalan             MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguageCatalan
	MicrosoftTokenizerLanguageChineseSimplified   MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguageChineseSimplified
	MicrosoftTokenizerLanguageChineseTraditional  MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguageChineseTraditional
	MicrosoftTokenizerLanguageCroatian            MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguageCroatian
	MicrosoftTokenizerLanguageCzech               MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguageCzech
	MicrosoftTokenizerLanguageDanish              MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguageDanish
	MicrosoftTokenizerLanguageDutch               MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguageDutch
	MicrosoftTokenizerLanguageEnglish             MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguageEnglish
	MicrosoftTokenizerLanguageFrench              MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguageFrench
	MicrosoftTokenizerLanguageGerman              MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguageGerman
	MicrosoftTokenizerLanguageGreek               MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguageGreek
	MicrosoftTokenizerLanguageGujarati            MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguageGujarati
	MicrosoftTokenizerLanguageHindi               MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguageHindi
	MicrosoftTokenizerLanguageIcelandic           MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguageIcelandic
	MicrosoftTokenizerLanguageIndonesian          MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguageIndonesian
	MicrosoftTokenizerLanguageItalian             MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguageItalian
	MicrosoftTokenizerLanguageJapanese            MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguageJapanese
	MicrosoftTokenizerLanguageKannada             MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguageKannada
	MicrosoftTokenizerLanguageKorean              MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguageKorean
	MicrosoftTokenizerLanguageMalay               MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguageMalay
	MicrosoftTokenizerLanguageMalayalam           MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguageMalayalam
	MicrosoftTokenizerLanguageMarathi             MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguageMarathi
	MicrosoftTokenizerLanguageNorwegianBokmaal    MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguageNorwegianBokmaal
	MicrosoftTokenizerLanguagePolish              MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguagePolish
	MicrosoftTokenizerLanguagePortuguese          MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguagePortuguese
	MicrosoftTokenizerLanguagePortugueseBrazilian MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguagePortugueseBrazilian
	MicrosoftTokenizerLanguagePunjabi             MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguagePunjabi
	MicrosoftTokenizerLanguageRomanian            MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguageRomanian
	MicrosoftTokenizerLanguageRussian             MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguageRussian
	MicrosoftTokenizerLanguageSerbianCyrillic     MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguageSerbianCyrillic
	MicrosoftTokenizerLanguageSerbianLatin        MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguageSerbianLatin
	MicrosoftTokenizerLanguageSlovenian           MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguageSlovenian
	MicrosoftTokenizerLanguageSpanish             MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguageSpanish
	MicrosoftTokenizerLanguageSwedish             MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguageSwedish
	MicrosoftTokenizerLanguageTamil               MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguageTamil
	MicrosoftTokenizerLanguageTelugu              MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguageTelugu
	MicrosoftTokenizerLanguageThai                MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguageThai
	MicrosoftTokenizerLanguageUkrainian           MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguageUkrainian
	MicrosoftTokenizerLanguageUrdu                MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguageUrdu
	MicrosoftTokenizerLanguageVietnamese          MicrosoftTokenizerLanguage = original.MicrosoftTokenizerLanguageVietnamese
)

type Mode = original.Mode

const (
	All Mode = original.All
	Any Mode = original.Any
)

type OdataType = original.OdataType

const (
	OdataTypeAnalyzer                             OdataType = original.OdataTypeAnalyzer
	OdataTypeMicrosoftAzureSearchCustomAnalyzer   OdataType = original.OdataTypeMicrosoftAzureSearchCustomAnalyzer
	OdataTypeMicrosoftAzureSearchPatternAnalyzer  OdataType = original.OdataTypeMicrosoftAzureSearchPatternAnalyzer
	OdataTypeMicrosoftAzureSearchStandardAnalyzer OdataType = original.OdataTypeMicrosoftAzureSearchStandardAnalyzer
	OdataTypeMicrosoftAzureSearchStopAnalyzer     OdataType = original.OdataTypeMicrosoftAzureSearchStopAnalyzer
)

type OdataTypeBasicCharFilter = original.OdataTypeBasicCharFilter

const (
	OdataTypeCharFilter                                   OdataTypeBasicCharFilter = original.OdataTypeCharFilter
	OdataTypeMicrosoftAzureSearchMappingCharFilter        OdataTypeBasicCharFilter = original.OdataTypeMicrosoftAzureSearchMappingCharFilter
	OdataTypeMicrosoftAzureSearchPatternReplaceCharFilter OdataTypeBasicCharFilter = original.OdataTypeMicrosoftAzureSearchPatternReplaceCharFilter
)

type OdataTypeBasicDataChangeDetectionPolicy = original.OdataTypeBasicDataChangeDetectionPolicy

const (
	OdataTypeDataChangeDetectionPolicy                              OdataTypeBasicDataChangeDetectionPolicy = original.OdataTypeDataChangeDetectionPolicy
	OdataTypeMicrosoftAzureSearchHighWaterMarkChangeDetectionPolicy OdataTypeBasicDataChangeDetectionPolicy = original.OdataTypeMicrosoftAzureSearchHighWaterMarkChangeDetectionPolicy
	OdataTypeMicrosoftAzureSearchSQLIntegratedChangeTrackingPolicy  OdataTypeBasicDataChangeDetectionPolicy = original.OdataTypeMicrosoftAzureSearchSQLIntegratedChangeTrackingPolicy
)

type OdataTypeBasicDataDeletionDetectionPolicy = original.OdataTypeBasicDataDeletionDetectionPolicy

const (
	OdataTypeDataDeletionDetectionPolicy                                 OdataTypeBasicDataDeletionDetectionPolicy = original.OdataTypeDataDeletionDetectionPolicy
	OdataTypeMicrosoftAzureSearchSoftDeleteColumnDeletionDetectionPolicy OdataTypeBasicDataDeletionDetectionPolicy = original.OdataTypeMicrosoftAzureSearchSoftDeleteColumnDeletionDetectionPolicy
)

type OdataTypeBasicTokenFilter = original.OdataTypeBasicTokenFilter

const (
	OdataTypeMicrosoftAzureSearchASCIIFoldingTokenFilter           OdataTypeBasicTokenFilter = original.OdataTypeMicrosoftAzureSearchASCIIFoldingTokenFilter
	OdataTypeMicrosoftAzureSearchCjkBigramTokenFilter              OdataTypeBasicTokenFilter = original.OdataTypeMicrosoftAzureSearchCjkBigramTokenFilter
	OdataTypeMicrosoftAzureSearchCommonGramTokenFilter             OdataTypeBasicTokenFilter = original.OdataTypeMicrosoftAzureSearchCommonGramTokenFilter
	OdataTypeMicrosoftAzureSearchDictionaryDecompounderTokenFilter OdataTypeBasicTokenFilter = original.OdataTypeMicrosoftAzureSearchDictionaryDecompounderTokenFilter
	OdataTypeMicrosoftAzureSearchEdgeNGramTokenFilter              OdataTypeBasicTokenFilter = original.OdataTypeMicrosoftAzureSearchEdgeNGramTokenFilter
	OdataTypeMicrosoftAzureSearchEdgeNGramTokenFilterV2            OdataTypeBasicTokenFilter = original.OdataTypeMicrosoftAzureSearchEdgeNGramTokenFilterV2
	OdataTypeMicrosoftAzureSearchElisionTokenFilter                OdataTypeBasicTokenFilter = original.OdataTypeMicrosoftAzureSearchElisionTokenFilter
	OdataTypeMicrosoftAzureSearchKeepTokenFilter                   OdataTypeBasicTokenFilter = original.OdataTypeMicrosoftAzureSearchKeepTokenFilter
	OdataTypeMicrosoftAzureSearchKeywordMarkerTokenFilter          OdataTypeBasicTokenFilter = original.OdataTypeMicrosoftAzureSearchKeywordMarkerTokenFilter
	OdataTypeMicrosoftAzureSearchLengthTokenFilter                 OdataTypeBasicTokenFilter = original.OdataTypeMicrosoftAzureSearchLengthTokenFilter
	OdataTypeMicrosoftAzureSearchLimitTokenFilter                  OdataTypeBasicTokenFilter = original.OdataTypeMicrosoftAzureSearchLimitTokenFilter
	OdataTypeMicrosoftAzureSearchNGramTokenFilter                  OdataTypeBasicTokenFilter = original.OdataTypeMicrosoftAzureSearchNGramTokenFilter
	OdataTypeMicrosoftAzureSearchNGramTokenFilterV2                OdataTypeBasicTokenFilter = original.OdataTypeMicrosoftAzureSearchNGramTokenFilterV2
	OdataTypeMicrosoftAzureSearchPatternCaptureTokenFilter         OdataTypeBasicTokenFilter = original.OdataTypeMicrosoftAzureSearchPatternCaptureTokenFilter
	OdataTypeMicrosoftAzureSearchPatternReplaceTokenFilter         OdataTypeBasicTokenFilter = original.OdataTypeMicrosoftAzureSearchPatternReplaceTokenFilter
	OdataTypeMicrosoftAzureSearchPhoneticTokenFilter               OdataTypeBasicTokenFilter = original.OdataTypeMicrosoftAzureSearchPhoneticTokenFilter
	OdataTypeMicrosoftAzureSearchShingleTokenFilter                OdataTypeBasicTokenFilter = original.OdataTypeMicrosoftAzureSearchShingleTokenFilter
	OdataTypeMicrosoftAzureSearchSnowballTokenFilter               OdataTypeBasicTokenFilter = original.OdataTypeMicrosoftAzureSearchSnowballTokenFilter
	OdataTypeMicrosoftAzureSearchStemmerOverrideTokenFilter        OdataTypeBasicTokenFilter = original.OdataTypeMicrosoftAzureSearchStemmerOverrideTokenFilter
	OdataTypeMicrosoftAzureSearchStemmerTokenFilter                OdataTypeBasicTokenFilter = original.OdataTypeMicrosoftAzureSearchStemmerTokenFilter
	OdataTypeMicrosoftAzureSearchStopwordsTokenFilter              OdataTypeBasicTokenFilter = original.OdataTypeMicrosoftAzureSearchStopwordsTokenFilter
	OdataTypeMicrosoftAzureSearchSynonymTokenFilter                OdataTypeBasicTokenFilter = original.OdataTypeMicrosoftAzureSearchSynonymTokenFilter
	OdataTypeMicrosoftAzureSearchTruncateTokenFilter               OdataTypeBasicTokenFilter = original.OdataTypeMicrosoftAzureSearchTruncateTokenFilter
	OdataTypeMicrosoftAzureSearchUniqueTokenFilter                 OdataTypeBasicTokenFilter = original.OdataTypeMicrosoftAzureSearchUniqueTokenFilter
	OdataTypeMicrosoftAzureSearchWordDelimiterTokenFilter          OdataTypeBasicTokenFilter = original.OdataTypeMicrosoftAzureSearchWordDelimiterTokenFilter
	OdataTypeTokenFilter                                           OdataTypeBasicTokenFilter = original.OdataTypeTokenFilter
)

type OdataTypeBasicTokenizer = original.OdataTypeBasicTokenizer

const (
	OdataTypeMicrosoftAzureSearchClassicTokenizer                   OdataTypeBasicTokenizer = original.OdataTypeMicrosoftAzureSearchClassicTokenizer
	OdataTypeMicrosoftAzureSearchEdgeNGramTokenizer                 OdataTypeBasicTokenizer = original.OdataTypeMicrosoftAzureSearchEdgeNGramTokenizer
	OdataTypeMicrosoftAzureSearchKeywordTokenizer                   OdataTypeBasicTokenizer = original.OdataTypeMicrosoftAzureSearchKeywordTokenizer
	OdataTypeMicrosoftAzureSearchKeywordTokenizerV2                 OdataTypeBasicTokenizer = original.OdataTypeMicrosoftAzureSearchKeywordTokenizerV2
	OdataTypeMicrosoftAzureSearchMicrosoftLanguageStemmingTokenizer OdataTypeBasicTokenizer = original.OdataTypeMicrosoftAzureSearchMicrosoftLanguageStemmingTokenizer
	OdataTypeMicrosoftAzureSearchMicrosoftLanguageTokenizer         OdataTypeBasicTokenizer = original.OdataTypeMicrosoftAzureSearchMicrosoftLanguageTokenizer
	OdataTypeMicrosoftAzureSearchNGramTokenizer                     OdataTypeBasicTokenizer = original.OdataTypeMicrosoftAzureSearchNGramTokenizer
	OdataTypeMicrosoftAzureSearchPathHierarchyTokenizer             OdataTypeBasicTokenizer = original.OdataTypeMicrosoftAzureSearchPathHierarchyTokenizer
	OdataTypeMicrosoftAzureSearchPathHierarchyTokenizerV2           OdataTypeBasicTokenizer = original.OdataTypeMicrosoftAzureSearchPathHierarchyTokenizerV2
	OdataTypeMicrosoftAzureSearchPatternTokenizer                   OdataTypeBasicTokenizer = original.OdataTypeMicrosoftAzureSearchPatternTokenizer
	OdataTypeMicrosoftAzureSearchStandardTokenizer                  OdataTypeBasicTokenizer = original.OdataTypeMicrosoftAzureSearchStandardTokenizer
	OdataTypeMicrosoftAzureSearchStandardTokenizerV2                OdataTypeBasicTokenizer = original.OdataTypeMicrosoftAzureSearchStandardTokenizerV2
	OdataTypeMicrosoftAzureSearchUaxURLEmailTokenizer               OdataTypeBasicTokenizer = original.OdataTypeMicrosoftAzureSearchUaxURLEmailTokenizer
	OdataTypeTokenizer                                              OdataTypeBasicTokenizer = original.OdataTypeTokenizer
)

type PhoneticEncoder = original.PhoneticEncoder

const (
	BeiderMorse     PhoneticEncoder = original.BeiderMorse
	Caverphone1     PhoneticEncoder = original.Caverphone1
	Caverphone2     PhoneticEncoder = original.Caverphone2
	Cologne         PhoneticEncoder = original.Cologne
	DoubleMetaphone PhoneticEncoder = original.DoubleMetaphone
	HaasePhonetik   PhoneticEncoder = original.HaasePhonetik
	KoelnerPhonetik PhoneticEncoder = original.KoelnerPhonetik
	Metaphone       PhoneticEncoder = original.Metaphone
	Nysiis          PhoneticEncoder = original.Nysiis
	RefinedSoundex  PhoneticEncoder = original.RefinedSoundex
	Soundex         PhoneticEncoder = original.Soundex
)

type QueryType = original.QueryType

const (
	Full   QueryType = original.Full
	Simple QueryType = original.Simple
)

type ScoringFunctionAggregation = original.ScoringFunctionAggregation

const (
	Average       ScoringFunctionAggregation = original.Average
	FirstMatching ScoringFunctionAggregation = original.FirstMatching
	Maximum       ScoringFunctionAggregation = original.Maximum
	Minimum       ScoringFunctionAggregation = original.Minimum
	Sum           ScoringFunctionAggregation = original.Sum
)

type ScoringFunctionInterpolation = original.ScoringFunctionInterpolation

const (
	Constant    ScoringFunctionInterpolation = original.Constant
	Linear      ScoringFunctionInterpolation = original.Linear
	Logarithmic ScoringFunctionInterpolation = original.Logarithmic
	Quadratic   ScoringFunctionInterpolation = original.Quadratic
)

type SnowballTokenFilterLanguage = original.SnowballTokenFilterLanguage

const (
	SnowballTokenFilterLanguageArmenian   SnowballTokenFilterLanguage = original.SnowballTokenFilterLanguageArmenian
	SnowballTokenFilterLanguageBasque     SnowballTokenFilterLanguage = original.SnowballTokenFilterLanguageBasque
	SnowballTokenFilterLanguageCatalan    SnowballTokenFilterLanguage = original.SnowballTokenFilterLanguageCatalan
	SnowballTokenFilterLanguageDanish     SnowballTokenFilterLanguage = original.SnowballTokenFilterLanguageDanish
	SnowballTokenFilterLanguageDutch      SnowballTokenFilterLanguage = original.SnowballTokenFilterLanguageDutch
	SnowballTokenFilterLanguageEnglish    SnowballTokenFilterLanguage = original.SnowballTokenFilterLanguageEnglish
	SnowballTokenFilterLanguageFinnish    SnowballTokenFilterLanguage = original.SnowballTokenFilterLanguageFinnish
	SnowballTokenFilterLanguageFrench     SnowballTokenFilterLanguage = original.SnowballTokenFilterLanguageFrench
	SnowballTokenFilterLanguageGerman     SnowballTokenFilterLanguage = original.SnowballTokenFilterLanguageGerman
	SnowballTokenFilterLanguageGerman2    SnowballTokenFilterLanguage = original.SnowballTokenFilterLanguageGerman2
	SnowballTokenFilterLanguageHungarian  SnowballTokenFilterLanguage = original.SnowballTokenFilterLanguageHungarian
	SnowballTokenFilterLanguageItalian    SnowballTokenFilterLanguage = original.SnowballTokenFilterLanguageItalian
	SnowballTokenFilterLanguageKp         SnowballTokenFilterLanguage = original.SnowballTokenFilterLanguageKp
	SnowballTokenFilterLanguageLovins     SnowballTokenFilterLanguage = original.SnowballTokenFilterLanguageLovins
	SnowballTokenFilterLanguageNorwegian  SnowballTokenFilterLanguage = original.SnowballTokenFilterLanguageNorwegian
	SnowballTokenFilterLanguagePorter     SnowballTokenFilterLanguage = original.SnowballTokenFilterLanguagePorter
	SnowballTokenFilterLanguagePortuguese SnowballTokenFilterLanguage = original.SnowballTokenFilterLanguagePortuguese
	SnowballTokenFilterLanguageRomanian   SnowballTokenFilterLanguage = original.SnowballTokenFilterLanguageRomanian
	SnowballTokenFilterLanguageRussian    SnowballTokenFilterLanguage = original.SnowballTokenFilterLanguageRussian
	SnowballTokenFilterLanguageSpanish    SnowballTokenFilterLanguage = original.SnowballTokenFilterLanguageSpanish
	SnowballTokenFilterLanguageSwedish    SnowballTokenFilterLanguage = original.SnowballTokenFilterLanguageSwedish
	SnowballTokenFilterLanguageTurkish    SnowballTokenFilterLanguage = original.SnowballTokenFilterLanguageTurkish
)

type StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguage

const (
	StemmerTokenFilterLanguageArabic            StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageArabic
	StemmerTokenFilterLanguageArmenian          StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageArmenian
	StemmerTokenFilterLanguageBasque            StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageBasque
	StemmerTokenFilterLanguageBrazilian         StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageBrazilian
	StemmerTokenFilterLanguageBulgarian         StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageBulgarian
	StemmerTokenFilterLanguageCatalan           StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageCatalan
	StemmerTokenFilterLanguageCzech             StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageCzech
	StemmerTokenFilterLanguageDanish            StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageDanish
	StemmerTokenFilterLanguageDutch             StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageDutch
	StemmerTokenFilterLanguageDutchKp           StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageDutchKp
	StemmerTokenFilterLanguageEnglish           StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageEnglish
	StemmerTokenFilterLanguageFinnish           StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageFinnish
	StemmerTokenFilterLanguageFrench            StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageFrench
	StemmerTokenFilterLanguageGalician          StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageGalician
	StemmerTokenFilterLanguageGerman            StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageGerman
	StemmerTokenFilterLanguageGerman2           StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageGerman2
	StemmerTokenFilterLanguageGreek             StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageGreek
	StemmerTokenFilterLanguageHindi             StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageHindi
	StemmerTokenFilterLanguageHungarian         StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageHungarian
	StemmerTokenFilterLanguageIndonesian        StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageIndonesian
	StemmerTokenFilterLanguageIrish             StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageIrish
	StemmerTokenFilterLanguageItalian           StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageItalian
	StemmerTokenFilterLanguageLatvian           StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageLatvian
	StemmerTokenFilterLanguageLightEnglish      StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageLightEnglish
	StemmerTokenFilterLanguageLightFinnish      StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageLightFinnish
	StemmerTokenFilterLanguageLightFrench       StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageLightFrench
	StemmerTokenFilterLanguageLightGerman       StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageLightGerman
	StemmerTokenFilterLanguageLightHungarian    StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageLightHungarian
	StemmerTokenFilterLanguageLightItalian      StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageLightItalian
	StemmerTokenFilterLanguageLightNorwegian    StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageLightNorwegian
	StemmerTokenFilterLanguageLightNynorsk      StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageLightNynorsk
	StemmerTokenFilterLanguageLightPortuguese   StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageLightPortuguese
	StemmerTokenFilterLanguageLightRussian      StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageLightRussian
	StemmerTokenFilterLanguageLightSpanish      StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageLightSpanish
	StemmerTokenFilterLanguageLightSwedish      StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageLightSwedish
	StemmerTokenFilterLanguageLovins            StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageLovins
	StemmerTokenFilterLanguageMinimalEnglish    StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageMinimalEnglish
	StemmerTokenFilterLanguageMinimalFrench     StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageMinimalFrench
	StemmerTokenFilterLanguageMinimalGalician   StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageMinimalGalician
	StemmerTokenFilterLanguageMinimalGerman     StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageMinimalGerman
	StemmerTokenFilterLanguageMinimalNorwegian  StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageMinimalNorwegian
	StemmerTokenFilterLanguageMinimalNynorsk    StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageMinimalNynorsk
	StemmerTokenFilterLanguageMinimalPortuguese StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageMinimalPortuguese
	StemmerTokenFilterLanguageNorwegian         StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageNorwegian
	StemmerTokenFilterLanguagePorter2           StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguagePorter2
	StemmerTokenFilterLanguagePortuguese        StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguagePortuguese
	StemmerTokenFilterLanguagePortugueseRslp    StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguagePortugueseRslp
	StemmerTokenFilterLanguagePossessiveEnglish StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguagePossessiveEnglish
	StemmerTokenFilterLanguageRomanian          StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageRomanian
	StemmerTokenFilterLanguageRussian           StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageRussian
	StemmerTokenFilterLanguageSorani            StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageSorani
	StemmerTokenFilterLanguageSpanish           StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageSpanish
	StemmerTokenFilterLanguageSwedish           StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageSwedish
	StemmerTokenFilterLanguageTurkish           StemmerTokenFilterLanguage = original.StemmerTokenFilterLanguageTurkish
)

type StopwordsList = original.StopwordsList

const (
	StopwordsListArabic     StopwordsList = original.StopwordsListArabic
	StopwordsListArmenian   StopwordsList = original.StopwordsListArmenian
	StopwordsListBasque     StopwordsList = original.StopwordsListBasque
	StopwordsListBrazilian  StopwordsList = original.StopwordsListBrazilian
	StopwordsListBulgarian  StopwordsList = original.StopwordsListBulgarian
	StopwordsListCatalan    StopwordsList = original.StopwordsListCatalan
	StopwordsListCzech      StopwordsList = original.StopwordsListCzech
	StopwordsListDanish     StopwordsList = original.StopwordsListDanish
	StopwordsListDutch      StopwordsList = original.StopwordsListDutch
	StopwordsListEnglish    StopwordsList = original.StopwordsListEnglish
	StopwordsListFinnish    StopwordsList = original.StopwordsListFinnish
	StopwordsListFrench     StopwordsList = original.StopwordsListFrench
	StopwordsListGalician   StopwordsList = original.StopwordsListGalician
	StopwordsListGerman     StopwordsList = original.StopwordsListGerman
	StopwordsListGreek      StopwordsList = original.StopwordsListGreek
	StopwordsListHindi      StopwordsList = original.StopwordsListHindi
	StopwordsListHungarian  StopwordsList = original.StopwordsListHungarian
	StopwordsListIndonesian StopwordsList = original.StopwordsListIndonesian
	StopwordsListIrish      StopwordsList = original.StopwordsListIrish
	StopwordsListItalian    StopwordsList = original.StopwordsListItalian
	StopwordsListLatvian    StopwordsList = original.StopwordsListLatvian
	StopwordsListNorwegian  StopwordsList = original.StopwordsListNorwegian
	StopwordsListPersian    StopwordsList = original.StopwordsListPersian
	StopwordsListPortuguese StopwordsList = original.StopwordsListPortuguese
	StopwordsListRomanian   StopwordsList = original.StopwordsListRomanian
	StopwordsListRussian    StopwordsList = original.StopwordsListRussian
	StopwordsListSorani     StopwordsList = original.StopwordsListSorani
	StopwordsListSpanish    StopwordsList = original.StopwordsListSpanish
	StopwordsListSwedish    StopwordsList = original.StopwordsListSwedish
	StopwordsListThai       StopwordsList = original.StopwordsListThai
	StopwordsListTurkish    StopwordsList = original.StopwordsListTurkish
)

type SuggesterSearchMode = original.SuggesterSearchMode

const (
	AnalyzingInfixMatching SuggesterSearchMode = original.AnalyzingInfixMatching
)

type TokenCharacterKind = original.TokenCharacterKind

const (
	Digit       TokenCharacterKind = original.Digit
	Letter      TokenCharacterKind = original.Letter
	Punctuation TokenCharacterKind = original.Punctuation
	Symbol      TokenCharacterKind = original.Symbol
	Whitespace  TokenCharacterKind = original.Whitespace
)

type Type = original.Type

const (
	TypeDistance        Type = original.TypeDistance
	TypeFreshness       Type = original.TypeFreshness
	TypeMagnitude       Type = original.TypeMagnitude
	TypeScoringFunction Type = original.TypeScoringFunction
	TypeTag             Type = original.TypeTag
)

type BasicAnalyzer = original.BasicAnalyzer
type Analyzer = original.Analyzer
type AnalyzeRequest = original.AnalyzeRequest
type AnalyzeResult = original.AnalyzeResult
type AnalyzerName = original.AnalyzerName
type ASCIIFoldingTokenFilter = original.ASCIIFoldingTokenFilter
type BasicCharFilter = original.BasicCharFilter
type CharFilter = original.CharFilter
type CharFilterName = original.CharFilterName
type CjkBigramTokenFilter = original.CjkBigramTokenFilter
type ClassicTokenizer = original.ClassicTokenizer
type CommonGramTokenFilter = original.CommonGramTokenFilter
type CorsOptions = original.CorsOptions
type CustomAnalyzer = original.CustomAnalyzer
type BasicDataChangeDetectionPolicy = original.BasicDataChangeDetectionPolicy
type DataChangeDetectionPolicy = original.DataChangeDetectionPolicy
type DataContainer = original.DataContainer
type BasicDataDeletionDetectionPolicy = original.BasicDataDeletionDetectionPolicy
type DataDeletionDetectionPolicy = original.DataDeletionDetectionPolicy
type DataSource = original.DataSource
type DataSourceCredentials = original.DataSourceCredentials
type DataSourceListResult = original.DataSourceListResult
type DataSourceType = original.DataSourceType
type DataType = original.DataType
type DictionaryDecompounderTokenFilter = original.DictionaryDecompounderTokenFilter
type DistanceScoringFunction = original.DistanceScoringFunction
type DistanceScoringParameters = original.DistanceScoringParameters
type DocumentIndexResult = original.DocumentIndexResult
type EdgeNGramTokenFilter = original.EdgeNGramTokenFilter
type EdgeNGramTokenFilterV2 = original.EdgeNGramTokenFilterV2
type EdgeNGramTokenizer = original.EdgeNGramTokenizer
type ElisionTokenFilter = original.ElisionTokenFilter
type Field = original.Field
type FieldMapping = original.FieldMapping
type FieldMappingFunction = original.FieldMappingFunction
type FreshnessScoringFunction = original.FreshnessScoringFunction
type FreshnessScoringParameters = original.FreshnessScoringParameters
type HighWaterMarkChangeDetectionPolicy = original.HighWaterMarkChangeDetectionPolicy
type Index = original.Index
type Indexer = original.Indexer
type IndexerExecutionInfo = original.IndexerExecutionInfo
type IndexerExecutionResult = original.IndexerExecutionResult
type IndexerListResult = original.IndexerListResult
type IndexGetStatisticsResult = original.IndexGetStatisticsResult
type IndexingParameters = original.IndexingParameters
type IndexingResult = original.IndexingResult
type IndexingSchedule = original.IndexingSchedule
type IndexListResult = original.IndexListResult
type Int64 = original.Int64
type ItemError = original.ItemError
type KeepTokenFilter = original.KeepTokenFilter
type KeywordMarkerTokenFilter = original.KeywordMarkerTokenFilter
type KeywordTokenizer = original.KeywordTokenizer
type KeywordTokenizerV2 = original.KeywordTokenizerV2
type LengthTokenFilter = original.LengthTokenFilter
type LimitTokenFilter = original.LimitTokenFilter
type MagnitudeScoringFunction = original.MagnitudeScoringFunction
type MagnitudeScoringParameters = original.MagnitudeScoringParameters
type MappingCharFilter = original.MappingCharFilter
type MicrosoftLanguageStemmingTokenizer = original.MicrosoftLanguageStemmingTokenizer
type MicrosoftLanguageTokenizer = original.MicrosoftLanguageTokenizer
type NGramTokenFilter = original.NGramTokenFilter
type NGramTokenFilterV2 = original.NGramTokenFilterV2
type NGramTokenizer = original.NGramTokenizer
type ParametersPayload = original.ParametersPayload
type PathHierarchyTokenizer = original.PathHierarchyTokenizer
type PathHierarchyTokenizerV2 = original.PathHierarchyTokenizerV2
type PatternAnalyzer = original.PatternAnalyzer
type PatternCaptureTokenFilter = original.PatternCaptureTokenFilter
type PatternReplaceCharFilter = original.PatternReplaceCharFilter
type PatternReplaceTokenFilter = original.PatternReplaceTokenFilter
type PatternTokenizer = original.PatternTokenizer
type PhoneticTokenFilter = original.PhoneticTokenFilter
type RegexFlags = original.RegexFlags
type BasicScoringFunction = original.BasicScoringFunction
type ScoringFunction = original.ScoringFunction
type ScoringProfile = original.ScoringProfile
type ShingleTokenFilter = original.ShingleTokenFilter
type SnowballTokenFilter = original.SnowballTokenFilter
type SoftDeleteColumnDeletionDetectionPolicy = original.SoftDeleteColumnDeletionDetectionPolicy
type SQLIntegratedChangeTrackingPolicy = original.SQLIntegratedChangeTrackingPolicy
type StandardAnalyzer = original.StandardAnalyzer
type StandardTokenizer = original.StandardTokenizer
type StandardTokenizerV2 = original.StandardTokenizerV2
type StemmerOverrideTokenFilter = original.StemmerOverrideTokenFilter
type StemmerTokenFilter = original.StemmerTokenFilter
type StopAnalyzer = original.StopAnalyzer
type StopwordsTokenFilter = original.StopwordsTokenFilter
type Suggester = original.Suggester
type SuggestParametersPayload = original.SuggestParametersPayload
type SynonymTokenFilter = original.SynonymTokenFilter
type TagScoringFunction = original.TagScoringFunction
type TagScoringParameters = original.TagScoringParameters
type TextWeights = original.TextWeights
type BasicTokenFilter = original.BasicTokenFilter
type TokenFilter = original.TokenFilter
type TokenFilterName = original.TokenFilterName
type TokenInfo = original.TokenInfo
type BasicTokenizer = original.BasicTokenizer
type Tokenizer = original.Tokenizer
type TokenizerName = original.TokenizerName
type TruncateTokenFilter = original.TruncateTokenFilter
type UaxURLEmailTokenizer = original.UaxURLEmailTokenizer
type UniqueTokenFilter = original.UniqueTokenFilter
type WordDelimiterTokenFilter = original.WordDelimiterTokenFilter

func NewDataSourcesClient() DataSourcesClient {
	return original.NewDataSourcesClient()
}
func NewDataSourcesClientWithBaseURI(baseURI string) DataSourcesClient {
	return original.NewDataSourcesClientWithBaseURI(baseURI)
}
func NewDocumentsProxyClient() DocumentsProxyClient {
	return original.NewDocumentsProxyClient()
}
func NewDocumentsProxyClientWithBaseURI(baseURI string) DocumentsProxyClient {
	return original.NewDocumentsProxyClientWithBaseURI(baseURI)
}
func NewIndexersClient() IndexersClient {
	return original.NewIndexersClient()
}
func NewIndexersClientWithBaseURI(baseURI string) IndexersClient {
	return original.NewIndexersClientWithBaseURI(baseURI)
}
func NewIndexesClient() IndexesClient {
	return original.NewIndexesClient()
}
func NewIndexesClientWithBaseURI(baseURI string) IndexesClient {
	return original.NewIndexesClientWithBaseURI(baseURI)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
func New() BaseClient {
	return original.New()
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}
