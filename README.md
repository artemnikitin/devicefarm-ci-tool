## Tool that helps more easier run your apps in AWS Device Farm
[![Go Report Card](https://goreportcard.com/badge/github.com/artemnikitin/devicefarm-ci-tool)](https://goreportcard.com/report/github.com/artemnikitin/devicefarm-ci-tool) [![FOSSA Status](https://app.fossa.io/api/projects/git%2Bhttps%3A%2F%2Fgithub.com%2Fartemnikitin%2Fdevicefarm-ci-tool.svg?type=shield)](https://app.fossa.io/projects/git%2Bhttps%3A%2F%2Fgithub.com%2Fartemnikitin%2Fdevicefarm-ci-tool?ref=badge_shield)
  [![Build Status](https://travis-ci.org/artemnikitin/devicefarm-ci-tool.svg?branch=master)](https://travis-ci.org/artemnikitin/devicefarm-ci-tool)     
#### Description
Did you try to run an app in AWS Device Farm via CLI or API? It was easy, right? Right now you can probably say ARN of your project after wake up in the middle of the night.

This tool helps to run apps in AWS Device Farm. You don't need to know ARN because it's for a machine and not for a human.

It's not a replacement for existing AWS CLI tools. It was created for a very specific purpose, to run tests in CI without a lot of configuration that required for existing solutions. It based on an assumption that all setup is already done. It means, that if you will specify an unexisted project name, then the tool will not create it.
#### AWS Credentials

Set environment variables     
```
export AWS_ACCESS_KEY_ID=<key>    
export AWS_SECRET_ACCESS_KEY=<secret>
```     

#### Download
Get it via    
``` 
go get -u github.com/artemnikitin/devicefarm-ci-tool 
``` 

Or download binaries for Windows, MacOS or Linux from [latest release](https://github.com/artemnikitin/devicefarm-ci-tool/releases/latest)   
   
#### Run
Required launch parameters:
- ```project``` name of a project.
- ```app``` path to an app.    
Example:
```
devicefarm-ci-tool -project name -app /path/to/my/app.apk
```
By default, "BUILTIN_FUZZ" tests will be run for your app.

Optional parameters:
- ```run``` name of test run. Overrides value of `runName` parameter from config.     
Example: `-run myName` 
- ```test``` path to tests. Overrides value of `testPackagePath` parameter from config.     
Example: `-test /path/to/my/testapp.apk` 
- ```devices``` name of device pool where app will be run.      
Example:`-devices my-device-pool`      
- ```config``` path to config in JSON format.      
Example:`-config /path/to/config.json`   
- ```wait``` wait for an end of test run. Useful for CI. Disabled by default.     
Example:`-wait`  
- ```checkEvery``` checks every X seconds for test run completion. Default value is 5 second.    
Example: `-checkEvery 15`
- ```ignoreUnavailableDevices``` allows to consider test runs as passed for runs where tests passes on several devices, but some of devices were unavailable.    
Example: `-ignoreUnavailableDevices`
- ```testType``` allows to specify test type via command line. Test type should be one of test types available on AWS Device Farm, see [ScheduleRunTest](http://docs.aws.amazon.com/devicefarm/latest/APIReference/API_ScheduleRunTest.html)       
Example: `-testType INSTRUMENTATION`

#### Configuration file
All parameters in the configuration file are optional. Configuration file is based on a syntax of [ScheduleRun](http://docs.aws.amazon.com/devicefarm/latest/APIReference/API_ScheduleRun.html) request.        

Example of config (for current master):    
```json
{
   "name": "string",
   "projectArn": "string",
   "projectName": "string",
   "appArn": "string",
   "appPath": "string",
   "devicePoolArn": "string",
   "devicePoolPath": "string",
   "testPackagePath": "string",
   "extraDataPackagePath": "string",
   "configuration": { 
      "auxiliaryApps": [ "string" ],
      "billingMethod": "METERED|UNMETERED",
      "extraDataPackageArn": "string",
      "locale": "string",
      "location": { 
         "latitude": 11.11,
         "longitude": 22.22
      },
      "networkProfileArn": "string",
      "radios": { 
         "bluetooth": true,
         "gps": false,
         "nfc": true,
         "wifi": false
      }
   },
   "executionConfiguration": { 
      "accountsCleanup": false,
      "appPackagesCleanup": true,
      "jobTimeoutMinutes": 111
   },
   "test": { 
      "filter": "string",
      "parameters": { 
         "string" : "string" 
      },
      "testPackageArn": "string",
      "type": "BUILTIN_FUZZ|BUILTIN_EXPLORER|APPIUM_JAVA_JUNIT|APPIUM_JAVA_TESTNG|APPIUM_PYTHON|APPIUM_WEB_JAVA_JUNIT|APPIUM_WEB_JAVA_TESTNG|APPIUM_WEB_PYTHON|CALABASH|INSTRUMENTATION|UIAUTOMATION|UIAUTOMATOR|XCTEST"
   }
}
```    
Example of old config (for release 1.2.1):    
```json
{
   "runName":"name",
   "devicePoolArn":"ARN",
   "devicePoolName":"Top Devices",
   "test":{
      "type":"BUILTIN_FUZZ|BUILTIN_EXPLORER|APPIUM_JAVA_JUNIT|APPIUM_JAVA_TESTNG|APPIUM_PYTHON|APPIUM_WEB_JAVA_JUNIT|APPIUM_WEB_JAVA_TESTNG|APPIUM_WEB_PYTHON|CALABASH|INSTRUMENTATION|UIAUTOMATION|UIAUTOMATOR|XCTEST",
      "testPackageArn":"string",
      "testPackagePath":"string",
      "filter":"string",
      "parameters":{
         "key":"value",
         "key":"value"
      }
   },
   "additionalData":{
      "extraDataPackageArn":"string",
      "extraDataPackagePath":"string",
      "networkProfileArn":"string",
      "locale":"string",
      "location":{
         "latitude":1.222,
         "longitude":1.222
      },
      "radios":{
         "wifi":"true|false",
         "bluetooth":"true|false",
         "nfc":"true|false",
         "gps":"true|false"
      },
      "auxiliaryApps":[
         "string1",
         "string2"
      ],
      "billingMethod":"METERED|UNMETERED"
   },
   "executionConfiguration":{ 
      "jobTimeoutMinutes":111,
      "accountsCleanup":false,
      "appPackagesCleanup":false
   }
}
```


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bhttps%3A%2F%2Fgithub.com%2Fartemnikitin%2Fdevicefarm-ci-tool.svg?type=large)](https://app.fossa.io/projects/git%2Bhttps%3A%2F%2Fgithub.com%2Fartemnikitin%2Fdevicefarm-ci-tool?ref=badge_large)
