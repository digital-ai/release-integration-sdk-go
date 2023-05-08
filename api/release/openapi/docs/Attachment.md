# Attachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**File** | Pointer to **map[string]interface{}** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**ExportFilename** | Pointer to **string** |  | [optional] 
**FileUri** | Pointer to **string** |  | [optional] 
**Placeholders** | Pointer to **[]string** |  | [optional] 
**TextFileNamesRegex** | Pointer to **string** |  | [optional] 
**ExcludeFileNamesRegex** | Pointer to **string** |  | [optional] 
**FileEncodings** | Pointer to **map[string]string** |  | [optional] 
**Checksum** | Pointer to **string** |  | [optional] 

## Methods

### NewAttachment

`func NewAttachment() *Attachment`

NewAttachment instantiates a new Attachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentWithDefaults

`func NewAttachmentWithDefaults() *Attachment`

NewAttachmentWithDefaults instantiates a new Attachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Attachment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Attachment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Attachment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Attachment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Attachment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Attachment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Attachment) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Attachment) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFile

`func (o *Attachment) GetFile() map[string]interface{}`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *Attachment) GetFileOk() (*map[string]interface{}, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *Attachment) SetFile(v map[string]interface{})`

SetFile sets File field to given value.

### HasFile

`func (o *Attachment) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetContentType

`func (o *Attachment) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *Attachment) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *Attachment) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *Attachment) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetExportFilename

`func (o *Attachment) GetExportFilename() string`

GetExportFilename returns the ExportFilename field if non-nil, zero value otherwise.

### GetExportFilenameOk

`func (o *Attachment) GetExportFilenameOk() (*string, bool)`

GetExportFilenameOk returns a tuple with the ExportFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportFilename

`func (o *Attachment) SetExportFilename(v string)`

SetExportFilename sets ExportFilename field to given value.

### HasExportFilename

`func (o *Attachment) HasExportFilename() bool`

HasExportFilename returns a boolean if a field has been set.

### GetFileUri

`func (o *Attachment) GetFileUri() string`

GetFileUri returns the FileUri field if non-nil, zero value otherwise.

### GetFileUriOk

`func (o *Attachment) GetFileUriOk() (*string, bool)`

GetFileUriOk returns a tuple with the FileUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUri

`func (o *Attachment) SetFileUri(v string)`

SetFileUri sets FileUri field to given value.

### HasFileUri

`func (o *Attachment) HasFileUri() bool`

HasFileUri returns a boolean if a field has been set.

### GetPlaceholders

`func (o *Attachment) GetPlaceholders() []string`

GetPlaceholders returns the Placeholders field if non-nil, zero value otherwise.

### GetPlaceholdersOk

`func (o *Attachment) GetPlaceholdersOk() (*[]string, bool)`

GetPlaceholdersOk returns a tuple with the Placeholders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholders

`func (o *Attachment) SetPlaceholders(v []string)`

SetPlaceholders sets Placeholders field to given value.

### HasPlaceholders

`func (o *Attachment) HasPlaceholders() bool`

HasPlaceholders returns a boolean if a field has been set.

### GetTextFileNamesRegex

`func (o *Attachment) GetTextFileNamesRegex() string`

GetTextFileNamesRegex returns the TextFileNamesRegex field if non-nil, zero value otherwise.

### GetTextFileNamesRegexOk

`func (o *Attachment) GetTextFileNamesRegexOk() (*string, bool)`

GetTextFileNamesRegexOk returns a tuple with the TextFileNamesRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextFileNamesRegex

`func (o *Attachment) SetTextFileNamesRegex(v string)`

SetTextFileNamesRegex sets TextFileNamesRegex field to given value.

### HasTextFileNamesRegex

`func (o *Attachment) HasTextFileNamesRegex() bool`

HasTextFileNamesRegex returns a boolean if a field has been set.

### GetExcludeFileNamesRegex

`func (o *Attachment) GetExcludeFileNamesRegex() string`

GetExcludeFileNamesRegex returns the ExcludeFileNamesRegex field if non-nil, zero value otherwise.

### GetExcludeFileNamesRegexOk

`func (o *Attachment) GetExcludeFileNamesRegexOk() (*string, bool)`

GetExcludeFileNamesRegexOk returns a tuple with the ExcludeFileNamesRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFileNamesRegex

`func (o *Attachment) SetExcludeFileNamesRegex(v string)`

SetExcludeFileNamesRegex sets ExcludeFileNamesRegex field to given value.

### HasExcludeFileNamesRegex

`func (o *Attachment) HasExcludeFileNamesRegex() bool`

HasExcludeFileNamesRegex returns a boolean if a field has been set.

### GetFileEncodings

`func (o *Attachment) GetFileEncodings() map[string]string`

GetFileEncodings returns the FileEncodings field if non-nil, zero value otherwise.

### GetFileEncodingsOk

`func (o *Attachment) GetFileEncodingsOk() (*map[string]string, bool)`

GetFileEncodingsOk returns a tuple with the FileEncodings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileEncodings

`func (o *Attachment) SetFileEncodings(v map[string]string)`

SetFileEncodings sets FileEncodings field to given value.

### HasFileEncodings

`func (o *Attachment) HasFileEncodings() bool`

HasFileEncodings returns a boolean if a field has been set.

### GetChecksum

`func (o *Attachment) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *Attachment) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *Attachment) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *Attachment) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


