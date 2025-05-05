#include "stdafx.h"
#include <mi.h>

void EnumerateAndPrintInstanceNames(MI_Session* miSession, 
                        _In_z_ const wchar_t* namespaceName, 
                        const wchar_t* className)
{
    MI_Operation miOperation = MI_OPERATION_NULL;

    MI_Session_EnumerateInstances(miSession,      // Session 
                                  0,              // Flags
                                  NULL,           // Options
                                  namespaceName,  // CIM Namespace
                                  className,      // Class name
                                  MI_FALSE,       // Retrieve only keys
                                  NULL,           // Callbacks
                                  &amp;miOperation);  // Operation

    MI_Result miResult = MI_RESULT_OK;
    MI_Boolean moreResults;
    const MI_Char* errorString = NULL;
    MI_Uint32 instanceCount = 0;
    MI_Instance *miInstance;
    const MI_Instance* errorDetails = NULL;

    do
    {
        //Note that each instance becomes invalid after getting the next instance in the loop or
        //after closing the operation. Call the MI_Instance_Clone function to use an instance 
        //past this. Be sure to MI_Instance_Delete to close the cloned instance when finished.

        miResult = MI_Operation_GetInstance(&amp;miOperation,                     // Operation 
                                            (const MI_Instance**)&amp;miInstance, // Instance
                                            &amp;moreResults,                     // More results?
                                            &amp;miResult,                        // Result
                                            &amp;errorString,                     // Error message
                                            &amp;errorDetails);                   // Completion details
        if (miResult != MI_RESULT_OK)
        {
            wprintf(L"MI_Operation_GetInstance failed. MI_RESULT: %ld\n", miResult);
            break;
        }
        //The following demonstrates using the instance just received.
        if (miInstance)  
        {
            MI_Value value;
            MI_Type type;
            MI_Uint32 flags;

            //Athough the Name property is shown here to demonstrate, you could substitute another property
            miResult = MI_Instance_GetElement(miInstance,  // Instance
                                              L"Name",     // Element (property) name
                                              &amp;value,      // Element value
                                              &amp;type,       // Element type
                                              &amp;flags,      // Flags
                                              NULL);       // Index
            if (miResult != MI_RESULT_OK)
            {
                wprintf(L"MI_Instance_GetElement failed. MI_RESULT: %ld)\n", miResult);
                return;
            }
            wprintf(L"Process Name: %s\n", value.string);

            instanceCount++;
        }

    } while (miResult == MI_RESULT_OK &amp;&amp; moreResults == MI_TRUE);

    if (miResult != MI_RESULT_OK)
    {
        wprintf(L"Operation failed: MI_Result=%ld, errorString=%s\n", 
                miResult, errorString);
    }
    else
    {
        wprintf(L"Operation succeeded. Number of instances = %u\n", instanceCount);
    }

    miResult = MI_Operation_Close(&amp;miOperation);
    if (miResult != MI_RESULT_OK)
    {
        wprintf(L"MI_Operation_Close failed. MI_RESULT: %ld\n", miResult);
    }
}

int _tmain(int argc, _TCHAR* argv[])
{
    MI_Result miResult = MI_RESULT_OK;
    MI_Application miApplication = MI_APPLICATION_NULL;
    MI_Session miSession = MI_SESSION_NULL;
    MI_Operation miOperation = MI_OPERATION_NULL;

    miResult = MI_Application_Initialize(0,                   // Flags - Must be 0
                                         NULL,                // Application ID
                                         NULL,                // Extended error
                                         &amp;miApplication);     // Application
    if (miResult != MI_RESULT_OK)
    {
        wprintf(L"MI_Application_Initialize failed. MI_RESULT: %ld\n", miResult);
        return -1;
    }

    miResult = MI_Application_NewSession(&amp;miApplication, // Application 
                                         L"WINRM",       // WimRM Protocol
                                         L"localhost",   // Machine destination
                                         NULL,           // Options
                                         NULL,           // Callbacks
                                         NULL,           // Extended error
                                         &amp;miSession);    // Session 
    if (miResult != MI_RESULT_OK)
    {
        wprintf(L"MI_Application_NewSession failed. MI_RESULT: %ld\n", miResult);
        return -1;
    }

    EnumerateAndPrintInstanceNames(&amp;miSession, L"root\\cimv2", L"Win32_Process");    

    miResult = MI_Session_Close(&amp;miSession, NULL, NULL);
    if (miResult != MI_RESULT_OK)
    {
        wprintf(L"MI_Session_Close failed. MI_Result: %ld\n", miResult);
        return -1;
    }

    miResult = MI_Application_Close(&amp;miApplication);
    if (miResult != MI_RESULT_OK) 
    {
        wprintf(L"MI_Application_Close failed. MI_RESULT: %ld\n", miResult);
        return -1;
    }

    return 0;
}