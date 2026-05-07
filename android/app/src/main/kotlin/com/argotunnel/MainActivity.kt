package com.argotunnel

import android.content.Intent
import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.compose.foundation.layout.Column
import androidx.compose.material3.Button
import androidx.compose.material3.Text
import androidx.compose.ui.platform.LocalContext

class MainActivity : ComponentActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContent { ArgoTunnelUI() }
    }
}

@androidx.compose.runtime.Composable
fun ArgoTunnelUI() {
    val context = LocalContext.current
    Column {
        Text("ArgoTunnel Ultimate – AI/Quantum Anti‑Censorship")
        Button(onClick = {
            context.startService(Intent(context, ArgoVpnService::class.java))
        }) { Text("Connect") }
        Button(onClick = {
            context.stopService(Intent(context, ArgoVpnService::class.java))
        }) { Text("Disconnect") }
    }
}
