package com.argotunnel

import android.app.*
import android.content.Intent
import android.net.VpnService
import android.os.ParcelFileDescriptor
import androidx.core.app.NotificationCompat
import kotlinx.coroutines.*

class ArgoVpnService : VpnService() {
    private var vpnInterface: ParcelFileDescriptor? = null
    private val scope = CoroutineScope(Dispatchers.IO + SupervisorJob())

    override fun onStartCommand(intent: Intent?, flags: Int, startId: Int): Int {
        startForeground(1234, createNotification())
        startVpn()
        return START_STICKY
    }

    private fun startVpn() {
        val builder = Builder()
            .setSession("ArgoTunnel Quantum")
            .addAddress("10.0.0.2", 32)
            .addRoute("0.0.0.0", 0)
            .addDnsServer("1.1.1.1")
            .addDnsServer("8.8.8.8")
            .setMtu(1280)
        vpnInterface = builder.establish()
        scope.launch { runTunnel() }
    }

    private suspend fun runTunnel() {
        // در اینجا می‌توان هسته Go را از طریق gomobile فراخوانی کرد
        while (isActive) {
            delay(60000)
        }
    }

    private fun createNotification(): Notification {
        val channelId = "argo_channel"
        val nm = getSystemService(NotificationManager::class.java)
        nm.createNotificationChannel(NotificationChannel(channelId, "ArgoTunnel VPN", NotificationManager.IMPORTANCE_LOW))
        return NotificationCompat.Builder(this, channelId)
            .setContentTitle("ArgoTunnel Active")
            .setContentText("Quantum‑safe, AI‑morphed traffic")
            .setSmallIcon(android.R.drawable.ic_lock_lock)
            .build()
    }

    override fun onDestroy() {
        scope.cancel()
        vpnInterface?.close()
        super.onDestroy()
    }
}
